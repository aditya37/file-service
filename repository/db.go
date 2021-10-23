package repository

import (
	"context"
	"io"

	"github.com/aditya37/file-service/model"
)

type DBReadWriter interface {
	io.Closer
	SaveFileInfo(ctx context.Context, data model.MstFile) (int64, error)
	GetUploadType(ctx context.Context) ([]*model.MstUploadType, error)
}
