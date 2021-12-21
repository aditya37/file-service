package service

import (
	"context"

	"github.com/aditya37/file-service/utils"
)

func (f *service) DetailFile(ctx context.Context, request DetailFileRequest) (DetailFileResponse, error) {
	resp, err := f.db.GetDetailFile(ctx, request.ObjectName)
	if err != nil {
		return DetailFileResponse{}, &utils.CustomError{
			InternalError: err.Error(),
			Description:   "Object Name not found",
			Code:          ErrCodeDataNotFound,
		}
	}
	attrs, err := f.storage.GetObject(ctx, resp.ObjectName)
	if err != nil {
		return DetailFileResponse{}, &utils.CustomError{
			InternalError: err.Error(),
			Description:   "Failed to read data",
			Code:          ErrCodeReadFile,
		}
	}
	return DetailFileResponse{
		Id:        resp.ID,
		CreatedAt: resp.CreatedAt,
		Object:    attrs.Name,
		Metadata: Metadata{
			MediaLink:   attrs.MediaLink,
			FileSize:    attrs.Size,
			ContentType: attrs.ContentType,
		},
	}, nil
}
