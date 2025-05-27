package fileupload

import (
	"context"
)

type FileUpload interface {
	Upload(ctx context.Context, data []byte) (string, error)
}
