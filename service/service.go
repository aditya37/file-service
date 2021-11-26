package service

import (
	"context"

	db "github.com/aditya37/file-service/repository"
	repo_storage "github.com/aditya37/file-service/repository/firebase"
)

type FileService interface {
	FileUpload(ctx context.Context, request FileUploadRequest) (FileUploadResponse, error)
	GetFiles(ctx context.Context, request GetFileRequest) (GetFilesResponse, error)
	DetailFile(ctx context.Context, request DetailFileRequest) (DetailFileResponse, error)
}
type service struct {
	storage repo_storage.FirebaseStorage
	db      db.DBReadWriter
}

func NewFileService(storage repo_storage.FirebaseStorage, db db.DBReadWriter) (FileService, error) {
	return &service{storage: storage, db: db}, nil
}
