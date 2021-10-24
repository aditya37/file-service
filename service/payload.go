package service

import (
	"mime/multipart"
	"time"
)

// Constant for custom error code

const (

	// dynamic error code
	// error code for wrong request
	ErrCodeWrongRequest = 101

	// Static use error code
	ErrCodeWrongFileFormat      = 002
	ErrCodeFailedWriteTemp      = 003
	ErrCodeFailedCreateTemp     = 004
	ErrCodeReadFile             = 005
	ErrCodeFileLarge            = 006
	ErrCodeFailedParseObject    = 007
	ErrCodeFailedUploadToServer = 100

	// Type upload handler
	UploadTypePhotoProfile = "PHOTO_PROFILE"
	UploadTypeContent      = "UPLOAD_CONTENT"
	UploadTypeDocument     = "DOCUMENT"
)

type (
	// /file/upload.....
	FileInfo struct {
		FileName    string         `json:"string,omitempty"`
		FileSize    int64          `json:"file_size,omitempty"`
		ContentType string         `json:"content_type,omitempty"`
		File        multipart.File `json:"file,omitempty"`
	}
	FileUploadRequest struct {
		UploadType string   `json:"upload_type,omitempty"`
		FileDetail FileInfo `json:"file_detail,omitempty"`
	}
	FileUploadResponse struct {
		Id         int64  `json:"id,omitempty"`
		MediaLink  string `json:"media_link,omitempty"`
		ObjectName string `json:"object_name,omitempty"`
		FileSize   int64  `json:"file_size,omitempty"`
	}
	ProcessUpload struct {
		FileType   string
		UploadType string
		Filename   string
		File       multipart.File
	}
	// Menampung data ke map upload type
	UploadType struct {
		Id           int64
		ObjectPrefix string
		UploadType   string
	}

	// Get File Response
	Metadata struct {
		MediaLink   string `json:"media_link,omitempty"`
		FileSize    int64  `json:"file_size,omitempty"`
		ContentType string `json:"content_type,omitempty"`
	}
	FileItems struct {
		Id           int64     `json:"id,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		IsDeleted    int       `json:"is_deleted,omitempty"`
		UploadType   string    `json:"upload_type,omitempty"`
		ObjectPrefix string    `json:"object_prefix,omitempty"`
		Metadata     Metadata  `json:"meta_data,omitempty"`
	}
	GetFilesResponse struct {
		Count     int64        `json:"count,omitempty"`
		FileItems []*FileItems `json:"file_items,omitempty"`
	}
)
