package server

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/aditya37/file-service/service"
	"github.com/gorilla/mux"
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

func decodeRequestUploadedFile(ctx context.Context, request *http.Request) (interface{}, error) {
	request.URL.Query().Add("page", "")
	request.URL.Query().Add("itemPerPage", "")

	queryPage, available := request.URL.Query()["page"]
	if !available {
		return nil, errors.New("Please add page in request")
	}
	queryItemPerPage, available := request.URL.Query()["itemPerPage"]
	if !available {
		return nil, errors.New("Please add item perpage in request")
	}

	page, _ := strconv.Atoi(queryPage[0])
	itemPerPage, _ := strconv.Atoi(queryItemPerPage[0])

	return service.GetFileRequest{
		Page:        page,
		ItemPerPage: itemPerPage,
	}, nil
}

func decodeDetailFileRequest(ctx context.Context, request *http.Request) (interface{}, error) {

	params := mux.Vars(request)
	object, ok := params["object"]
	if !ok {
		return nil, errors.New("Object not found")
	}

	return service.DetailFileRequest{
		ObjectName: object,
	}, nil
}
