package server

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/aditya37/file-service/service"
)

func getContentType(source multipart.File) (string, error) {
	readBuffer := make([]byte, 1024)
	_, err := source.Read(readBuffer)
	if err != nil {
		return "", err
	}

	// get content type from buffer
	contentType := http.DetectContentType(readBuffer)
	return contentType, nil
}

// decodeRequestFileUpload
func decodeRequestFileUpload(ctx context.Context, request *http.Request) (interface{}, error) {
	var req service.FileUploadRequest

	uploadType := request.FormValue("upload_type")
	file, handler, err := request.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	src, err := handler.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	contentType, err := getContentType(src)
	if err != nil {
		return nil, err
	}

	req = service.FileUploadRequest{
		FileDetail: service.FileInfo{
			FileSize:    handler.Size,
			FileName:    handler.Filename,
			ContentType: contentType,
			File:        file,
		},
		UploadType: uploadType,
	}
	return req, nil
}
