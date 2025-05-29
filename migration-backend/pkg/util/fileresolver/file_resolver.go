package fileresolver

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type IPFSPath string

func (s IPFSPath) IPFSID() string {
	r, _ := strings.CutPrefix(string(s), "ipfs://")
	return r
}

func IsIPFSPath(path string) (*IPFSPath, bool) {
	if strings.HasPrefix(path, "ipfs://") {
		i := IPFSPath(path)
		return &i, true
	}
	return nil, false
}

type ArweavePath string

func (s ArweavePath) ArweaveID() string {
	r, _ := strings.CutPrefix(string(s), "ar://")
	return r
}

func IsArweavePath(path string) (*ArweavePath, bool) {
	if strings.HasPrefix(path, "ar://") {
		a := ArweavePath(path)
		return &a, true
	}
	return nil, false
}

type HTTPPath url.URL

func IsHTTPPath(path string) (*HTTPPath, bool) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		u, err := url.Parse(path)
		if err != nil {
			return nil, true
		}
		return (*HTTPPath)(u), false
	}
	return nil, false
}

type FileResolver interface {
	Resolve(ctx context.Context, path string) ([]byte, error)
}

type fileResolver struct {
	httpClient *http.Client

	ipfsHTTPBaseURL    string
	arweaveHTTPBaseURL string
}

func MakeFileResolver(
	httpClient *http.Client,
	ipfsHTTPBaseURL string,
	arweaveHTTPBaseURL string,
) FileResolver {
	return &fileResolver{
		httpClient:         httpClient,
		ipfsHTTPBaseURL:    ipfsHTTPBaseURL,
		arweaveHTTPBaseURL: arweaveHTTPBaseURL,
	}
}

func (r *fileResolver) Resolve(ctx context.Context, path string) ([]byte, error) {
	if p, isHTTP := IsHTTPPath(path); isHTTP {
		return r.resolveUrl(ctx, p)
	}

	if p, isIPFS := IsIPFSPath(path); isIPFS {
		return r.resolveIPFS(ctx, p)
	}

	if p, isArweave := IsArweavePath(path); isArweave {
		return r.resolveArweave(ctx, p)
	}

	return r.resolveBase(ctx, path)
}

func (r *fileResolver) resolveUrl(ctx context.Context, path *HTTPPath) ([]byte, error) {
	fmt.Println((*url.URL)(path).String())
	request, err := http.NewRequest("GET", (*url.URL)(path).String(), nil)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	resp, err := r.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func (r *fileResolver) resolveIPFS(ctx context.Context, path *IPFSPath) ([]byte, error) {
	u, err := url.Parse(r.ipfsHTTPBaseURL)
	if err != nil {
		return nil, err
	}
	u = u.JoinPath(path.IPFSID())
	return r.resolveUrl(ctx, (*HTTPPath)(u))
}

func (r *fileResolver) resolveArweave(ctx context.Context, path *ArweavePath) ([]byte, error) {
	u, err := url.Parse(r.arweaveHTTPBaseURL)
	if err != nil {
		return nil, err
	}
	u = u.JoinPath(path.ArweaveID())
	return r.resolveUrl(ctx, (*HTTPPath)(u))
}

func (r *fileResolver) resolveBase(_ context.Context, path string) ([]byte, error) {
	return os.ReadFile(path)
}
