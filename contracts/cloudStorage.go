package contracts

import (
	"context"
	"mime/multipart"
)

type ICloudStorageRepo interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, name string) (string, error)
	DeleteFile(ctx context.Context, fileUrl string) error
	Connect() error
}
