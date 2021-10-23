package endpoint

import (
	"context"

	"github.com/aditya37/file-service/service"
	"github.com/go-kit/kit/endpoint"
)

type FileServiceEndpoint struct {
	FileUploadEndpoint endpoint.Endpoint
}

func NewFileServiceEndpoint(srv service.FileService) FileServiceEndpoint {
	var fileUploadEndpoint endpoint.Endpoint
	{
		fileUploadEndpoint = MakeFileUploadEndpoint(srv)
	}
	return FileServiceEndpoint{
		FileUploadEndpoint: fileUploadEndpoint,
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
