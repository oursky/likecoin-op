package fileupload

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfsnode/data/builder"
	"github.com/ipld/go-car/v2/blockstore"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	delegationcore "github.com/web3-storage/go-ucanto/core/delegation"
	"github.com/web3-storage/go-ucanto/did"
	"github.com/web3-storage/go-ucanto/principal"
	"github.com/web3-storage/go-ucanto/principal/ed25519/signer"
	"github.com/web3-storage/go-w3up/capability/storeadd"
	"github.com/web3-storage/go-w3up/capability/uploadadd"
	"github.com/web3-storage/go-w3up/client"
	"github.com/web3-storage/go-w3up/delegation"
)

type w3client struct {
	logger *slog.Logger

	spaceDID   did.DID
	issuer     principal.Signer
	delegation delegationcore.Delegation
	httpClient *http.Client
}

func MakeS3Client(
	logger *slog.Logger,
	w3DID string,
	w3DIDKey string,
	w3UcanPath string,
	httpClient *http.Client,
) (FileUpload, error) {
	space, err := did.Parse(w3DID)
	if err != nil {
		return nil, err
	}
	issuer, err := signer.Parse(w3DIDKey)
	if err != nil {
		return nil, err
	}
	prfbytes, err := os.ReadFile(w3UcanPath)
	if err != nil {
		return nil, err
	}
	delegation, err := delegation.ExtractProof(prfbytes)
	if err != nil {
		return nil, err
	}

	return &w3client{
		logger:     logger.WithGroup("w3client"),
		spaceDID:   space,
		issuer:     issuer,
		delegation: delegation,
		httpClient: httpClient,
	}, nil
}

func (s *w3client) Upload(c context.Context, fileData []byte) (string, error) {
	logger := s.logger.WithGroup("Upload")
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		return "", err
	}
	logger = logger.With("tempFile", tempFile.Name())
	logger.Debug("tempFile created")
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write(fileData)
	if err != nil {
		return "", err
	}

	// https://github.com/ipld/go-car/blob/master/cmd/car/create.go

	hasher, err := multihash.GetHasher(multihash.SHA2_256)
	if err != nil {
		return "", err
	}
	digest := hasher.Sum([]byte{})
	hash, err := multihash.Encode(digest, multihash.SHA2_256)
	if err != nil {
		return "", err
	}
	proxyRoot := cid.NewCidV1(uint64(multicodec.DagPb), hash)
	logger = logger.With("proxyRoot", proxyRoot.String())

	tempCarFile, err := os.CreateTemp("", fmt.Sprintf("%s.*.car", proxyRoot.String()))
	if err != nil {
		return "", err
	}
	logger = logger.With("tempCarFile", tempCarFile.Name())
	logger.Debug("tempCarFile created")
	defer os.Remove(tempCarFile.Name())

	cdest, err := blockstore.OpenReadWrite(tempCarFile.Name(), []cid.Cid{proxyRoot})
	if err != nil {
		return "", err
	}

	// Write the unixfs blocks into the store.
	rootCid, err := s.writeFiles(c, true, cdest, tempFile.Name())
	if err != nil {
		return "", err
	}

	logger = logger.With("rootCid", rootCid.String())
	rootLink := cidlink.Link{Cid: rootCid}

	if err := cdest.Finalize(); err != nil {
		return "", err
	}

	// https://docs.storacha.network/go-w3up/
	carData, err := os.ReadFile(tempCarFile.Name())
	if err != nil {
		return "", err
	}

	mh, err := multihash.Sum(carData, multihash.SHA2_256, -1)
	if err != nil {
		return "", err
	}
	shardCid := cid.NewCidV1(0x0202, mh)
	logger = logger.With("shardCid", shardCid.String())
	// generate the CID for the CAR
	shardLink := cidlink.Link{Cid: shardCid}

	logger.Info("Instantiating shard upload...")
	storeAddRcpt, err := client.StoreAdd(
		s.issuer,
		s.spaceDID,
		&storeadd.Caveat{Link: shardLink, Size: uint64(len(carData))},
		client.WithProof(s.delegation),
	)

	if err != nil {
		return "", err
	}

	storeAddRcptFailure := storeAddRcpt.Out().Error()

	if storeAddRcptFailure != nil {
		return "", errors.New(storeAddRcptFailure.Message)
	}

	// "upload" means it needs to be uploaded, "done" means it is already done!
	if storeAddRcpt.Out().Ok().Status == "upload" {
		logger := logger.With("uploadUrl", *storeAddRcpt.Out().Ok().Url)
		logger.Info("Uploading shard...")
		hr, err := http.NewRequest("PUT", *storeAddRcpt.Out().Ok().Url, bytes.NewReader(carData))
		if err != nil {
			return "", err
		}

		hdr := map[string][]string{}
		for k, v := range storeAddRcpt.Out().Ok().Headers.Values {
			hdr[k] = []string{v}
		}

		hr.Header = hdr
		hr.ContentLength = int64(len(carData))
		res, err := s.httpClient.Do(hr)
		if err != nil {
			return "", err
		}
		res.Body.Close()
	}

	logger.Info("Attempt to registering an upload to w3store ipfs...")
	uploadAddRcpt, err := client.UploadAdd(
		s.issuer,
		s.spaceDID,
		&uploadadd.Caveat{
			Root:   rootLink,
			Shards: []ipld.Link{shardLink},
		},
		client.WithProof(s.delegation),
	)

	if err != nil {
		return "", nil
	}

	uploadAddRcptFailure := uploadAddRcpt.Out().Error()

	if uploadAddRcptFailure != nil {
		return "", errors.New(uploadAddRcptFailure.Message)
	}

	return s.makeUrl(rootLink), nil
}

func (s *w3client) makeUrl(link cidlink.Link) string {
	return fmt.Sprintf("ipfs://%s", link.String())
}

// https://github.com/ipld/go-car/blob/3e130683a0b7e047b069184952cc271d0f7e4342/cmd/car/create.go#L78
func (s *w3client) writeFiles(ctx context.Context, noWrap bool, bs *blockstore.ReadWrite, paths ...string) (cid.Cid, error) {
	ls := cidlink.DefaultLinkSystem()
	ls.TrustedStorage = true
	ls.StorageReadOpener = func(_ ipld.LinkContext, l ipld.Link) (io.Reader, error) {
		cl, ok := l.(cidlink.Link)
		if !ok {
			return nil, fmt.Errorf("not a cidlink")
		}
		blk, err := bs.Get(ctx, cl.Cid)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(blk.RawData()), nil
	}
	ls.StorageWriteOpener = func(_ ipld.LinkContext) (io.Writer, ipld.BlockWriteCommitter, error) {
		buf := bytes.NewBuffer(nil)
		return buf, func(l ipld.Link) error {
			cl, ok := l.(cidlink.Link)
			if !ok {
				return fmt.Errorf("not a cidlink")
			}
			blk, err := blocks.NewBlockWithCid(buf.Bytes(), cl.Cid)
			if err != nil {
				return err
			}
			bs.Put(ctx, blk)
			return nil
		}, nil
	}

	topLevel := make([]dagpb.PBLink, 0, len(paths))
	for _, p := range paths {
		l, size, err := builder.BuildUnixFSRecursive(p, &ls)
		if err != nil {
			return cid.Undef, err
		}
		if noWrap {
			rcl, ok := l.(cidlink.Link)
			if !ok {
				return cid.Undef, fmt.Errorf("could not interpret %s", l)
			}
			return rcl.Cid, nil
		}
		name := path.Base(p)
		entry, err := builder.BuildUnixFSDirectoryEntry(name, int64(size), l)
		if err != nil {
			return cid.Undef, err
		}
		topLevel = append(topLevel, entry)
	}

	// make a directory for the file(s).

	root, _, err := builder.BuildUnixFSDirectory(topLevel, &ls)
	if err != nil {
		return cid.Undef, nil
	}
	rcl, ok := root.(cidlink.Link)
	if !ok {
		return cid.Undef, fmt.Errorf("could not interpret %s", root)
	}

	return rcl.Cid, nil
}
