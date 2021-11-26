package endpoint

import (
	"context"

	"github.com/aditya37/file-service/service"
	"github.com/go-kit/kit/endpoint"
)

type FileServiceEndpoint struct {
	FileUploadEndpoint    endpoint.Endpoint
	UploadedFilesEndpoint endpoint.Endpoint
}

func NewFileServiceEndpoint(srv service.FileService) FileServiceEndpoint {
	var fileUploadEndpoint endpoint.Endpoint
	{
		fileUploadEndpoint = MakeFileUploadEndpoint(srv)
	}
	var uploadedFilesEndpoint endpoint.Endpoint
	{
		uploadedFilesEndpoint = MakeUploadedFilesEndpoint(srv)
	}
	return FileServiceEndpoint{
		FileUploadEndpoint:    fileUploadEndpoint,
		UploadedFilesEndpoint: uploadedFilesEndpoint,
	}
}

func MakeFileUploadEndpoint(srv service.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.FileUploadRequest)
		resp, err := srv.FileUpload(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func MakeUploadedFilesEndpoint(srv service.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.GetFileRequest)
		resp, err := srv.GetFiles(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
