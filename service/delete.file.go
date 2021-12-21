package service

import (
	"context"

	"github.com/aditya37/file-service/utils"
)

func (f *service) DeleteFile(ctx context.Context, request DeleteFileRequest) (DeleteFileResponse, error) {
	obj, err := f.db.GetDetailFile(ctx, request.ObjectName)
	if err != nil {
		return DeleteFileResponse{}, &utils.CustomError{
			InternalError: err.Error(),
			Description:   err.Error(),
			Code:          ErrCodeDataNotFound,
		}
	}

	// soft delete
	if err := f.db.DeleteUploadedFile(ctx, obj.ObjectName); err != nil {
		return DeleteFileResponse{}, &utils.CustomError{
			InternalError: err.Error(),
			Description:   err.Error(),
			Code:          500,
		}
	}

	// firebase delete obj
	if err := f.storage.DeleteObject(ctx, obj.ObjectName); err != nil {
		return DeleteFileResponse{}, &utils.CustomError{
			InternalError: err.Error(),
			Description:   err.Error(),
			Code:          500,
		}
	}
	return DeleteFileResponse{
		Message: "success",
	}, nil
}
