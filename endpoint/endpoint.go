package endpoint

import (
	"context"

	"github.com/aditya37/file-service/service"
	"github.com/go-kit/kit/endpoint"
)

type FileServiceEndpoint struct {
	FileUploadEndpoint    endpoint.Endpoint
	UploadedFilesEndpoint endpoint.Endpoint
	DetailFile            endpoint.Endpoint
	DeleteFile            endpoint.Endpoint
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
	var detailFile endpoint.Endpoint
	{
		detailFile = MakeDetailFile(srv)
	}
	var deleteFile endpoint.Endpoint
	{
		deleteFile = MakeDeleteFile(srv)
	}
	return FileServiceEndpoint{
		FileUploadEndpoint:    fileUploadEndpoint,
		UploadedFilesEndpoint: uploadedFilesEndpoint,
		DetailFile:            detailFile,
		DeleteFile:            deleteFile,
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

func MakeDetailFile(srv service.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.DetailFileRequest)
		resp, err := srv.DetailFile(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func MakeDeleteFile(srv service.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.DeleteFileRequest)
		resp, err := srv.DeleteFile(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
