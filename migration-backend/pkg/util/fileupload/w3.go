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
	"github.com/ipld/go-car/v2"
	"github.com/ipld/go-car/v2/blockstore"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/likecoin/like-migration-backend/pkg/util/fileupload/w3/filecoinoffer"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"github.com/storacha/go-piece/pkg/digest"
	"github.com/storacha/go-piece/pkg/piece"
	"github.com/web3-storage/go-ucanto/client"
	delegationcore "github.com/web3-storage/go-ucanto/core/delegation"
	"github.com/web3-storage/go-ucanto/core/invocation"
	"github.com/web3-storage/go-ucanto/did"
	"github.com/web3-storage/go-ucanto/principal"
	"github.com/web3-storage/go-ucanto/principal/ed25519/signer"
	"github.com/web3-storage/go-w3up/capability/storeadd"
	"github.com/web3-storage/go-w3up/capability/uploadadd"
	w3upclient "github.com/web3-storage/go-w3up/client"
	"github.com/web3-storage/go-w3up/delegation"
)

type w3client struct {
	dryRun    bool
	exportCar bool

	logger     *slog.Logger
	httpClient *http.Client

	spaceDID   did.DID
	issuer     principal.Signer
	delegation delegationcore.Delegation
}

func MakeW3Client(
	logger *slog.Logger,
	httpClient *http.Client,
	w3DID string,
	w3DIDKey string,
	w3UcanPath string,
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
		dryRun:     false,
		exportCar:  true,
		logger:     logger.WithGroup("w3client"),
		httpClient: httpClient,
		spaceDID:   space,
		issuer:     issuer,
		delegation: delegation,
	}, nil
}

// Check dry run
// go run ./cmd/cli fileupload upload test-images/1.webp
// go run ./cmd/cli image convert test-images/1.webp
// go run ./cmd/cli fileupload upload test-images/1.gif
// go run ./cmd/cli fileupload upload bafybeihhpejs7u4bpjbd3gwzultmoxap4yuzozblwyhem2x2aazpazvw2q.webp
// go run ./cmd/cli fileupload upload bafybeihhpejs7u4bpjbd3gwzultmoxap4yuzozblwyhem2x2aazpazvw2q.gif
// go run ./cmd/cli fileupload upload animated-webp-supported.webp
// go run ./cmd/cli fileupload upload animated-webp-supported.gif
// go run ./cmd/cli fileupload upload 1.sm.webp
// go run ./cmd/cli fileupload upload 1.sm.gif
// go run ./cmd/cli fileupload upload file_example_WEBP_1500kB.webp
// go run ./cmd/cli image convert file_example_WEBP_1500kB.webp
// go run ./cmd/cli fileupload upload file_example_WEBP_1500kB.gif
func (s *w3client) Upload(c context.Context, fileData []byte) (string, error) {
	logger := s.logger.WithGroup("Upload")

	rootCid, carData, err := s.makeCarFromBytes(c, fileData)
	if err != nil {
		return "", err
	}
	logger = logger.With("rootCid", rootCid.String())
	rootLink := cidlink.Link{Cid: rootCid}

	if s.exportCar {
		exportedCarFile, err := os.OpenFile("test-images/export.car", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}
		_, err = exportedCarFile.Write(carData)
		if err != nil {
			return "", err
		}
		exportedCarFile.Close()
	}

	// https://docs.storacha.network/go-w3up/
	mh, err := multihash.Sum(carData, multihash.SHA2_256, -1)
	if err != nil {
		return "", err
	}
	shardCid := cid.NewCidV1(0x0202, mh)
	logger = logger.With("shardCid", shardCid.String())
	// generate the CID for the CAR
	shardLink := cidlink.Link{Cid: shardCid}

	// filecoin content
	shardContentCid := cid.NewCidV1(digest.FR32_SHA256_TRUNC254_PADDED_BINARY_TREE_CODE, mh)
	shardContentLink := cidlink.Link{Cid: shardContentCid}

	// filecoin piece
	pieceH, err := multihash.Encode(carData, digest.FR32_SHA256_TRUNC254_PADDED_BINARY_TREE_CODE)
	if err != nil {
		return "", err
	}
	pieceMH, err := multihash.Cast(pieceH)
	if err != nil {
		return "", err
	}
	pd, err := digest.NewPieceDigest(pieceMH)
	if err != nil {
		return "", err
	}
	pieceLink := piece.FromPieceDigest(pd)

	logger.Info("Instantiating shard upload...")

	if s.dryRun {
		return s.makeUrl(rootLink), nil
	}

	storeAddRcpt, err := w3upclient.StoreAdd(
		s.issuer,
		s.spaceDID,
		&storeadd.Caveat{Link: shardLink, Size: uint64(len(carData))},
		w3upclient.WithProof(s.delegation),
	)

	if err != nil {
		return "", err
	}

	storeAddRcptFailure := storeAddRcpt.Out().Error()

	if storeAddRcptFailure != nil {
		return "", errors.New(storeAddRcptFailure.Message)
	}
	storeAddRecptSuccess := storeAddRcpt.Out().Ok()

	// "upload" means it needs to be uploaded, "done" means it is already done!
	if storeAddRecptSuccess.Status == "upload" {
		logger.Info("Uploading shard...")
		hr, err := http.NewRequest("PUT", *storeAddRecptSuccess.Url, bytes.NewReader(carData))
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
		defer res.Body.Close()
	}

	logger.Info("Attempt to registering an upload to w3store ipfs...")
	uploadAddRcpt, err := w3upclient.UploadAdd(
		s.issuer,
		s.spaceDID,
		&uploadadd.Caveat{
			Root:   rootLink,
			Shards: []ipld.Link{shardLink},
		},
		w3upclient.WithProof(s.delegation),
	)

	if err != nil {
		return "", nil
	}

	uploadAddRcptFailure := uploadAddRcpt.Out().Error()

	if uploadAddRcptFailure != nil {
		return "", errors.New(uploadAddRcptFailure.Message)
	}

	// https://github.com/storacha/w3up/blob/d02da30f51dfb79cfc9ce63f0b3a28aa9ea2345e/packages/upload-client/src/index.js#L175
	issuedInvocation, err := invocation.Invoke(
		s.issuer,
		w3upclient.DefaultConnection.ID(),
		filecoinoffer.NewCapability(s.spaceDID, &filecoinoffer.Caveat{
			Content: shardContentLink,
			Piece:   pieceLink,
		}),
		delegationcore.WithProofs([]delegationcore.Delegation{
			s.delegation,
		}),
	)
	if err != nil {
		return "", err
	}
	logger.Info("Invoking filecoin offer...")
	_, err = client.Execute([]invocation.Invocation{issuedInvocation}, w3upclient.DefaultConnection)
	if err != nil {
		return "", err
	}

	return s.makeUrl(rootLink), nil
}

func (s *w3client) makeUrl(link cidlink.Link) string {
	return fmt.Sprintf("ipfs://%s", link.String())
}

func (s *w3client) makeCarFromBytes(ctx context.Context, fileData []byte) (cid.Cid, []byte, error) {
	// Put file data in file system to be used by writeFiles call
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		return cid.Cid{}, nil, err
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write(fileData)
	if err != nil {
		return cid.Cid{}, nil, err
	}

	// The make car logic
	// ref: https://github.com/ipld/go-car/blob/master/cmd/car/create.go
	hasher, err := multihash.GetHasher(multihash.SHA2_256)
	if err != nil {
		return cid.Cid{}, nil, err
	}

	digest := hasher.Sum([]byte{})
	hash, err := multihash.Encode(digest, multihash.SHA2_256)
	if err != nil {
		return cid.Cid{}, nil, err
	}
	proxyRootCid := cid.NewCidV1(uint64(multicodec.DagPb), hash)

	tempCarFile, err := os.CreateTemp("", fmt.Sprintf("%s.*.car", proxyRootCid.String()))
	if err != nil {
		return cid.Cid{}, nil, err
	}
	defer os.Remove(tempCarFile.Name())

	carBlockStore, err := blockstore.OpenReadWrite(tempCarFile.Name(), []cid.Cid{proxyRootCid}, blockstore.WriteAsCarV1(true))
	if err != nil {
		return cid.Cid{}, nil, err
	}

	// Write the unixfs blocks into the store.
	rootCid, err := s.writeFiles(ctx, true, carBlockStore, tempFile.Name())
	if err != nil {
		return cid.Cid{}, nil, err
	}

	if err := carBlockStore.Finalize(); err != nil {
		return cid.Cid{}, nil, err
	}

	err = car.ReplaceRootsInFile(tempCarFile.Name(), []cid.Cid{rootCid})
	if err != nil {
		return cid.Cid{}, nil, err
	}

	carData, err := os.ReadFile(tempCarFile.Name())
	if err != nil {
		return cid.Cid{}, nil, err
	}

	return rootCid, carData, nil
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
