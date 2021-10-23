package service

import (
	"context"

	repo_storage "github.com/aditya37/file-service/repository/firebase"
)

type FileService interface {
	FileUpload(ctx context.Context, request FileUploadRequest) (FileUploadResponse, error)
}
type service struct {
	storage repo_storage.FirebaseStorage
}

func NewFileService(storage repo_storage.FirebaseStorage) (FileService, error) {
	return &service{storage: storage}, nil
}
