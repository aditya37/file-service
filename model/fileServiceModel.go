package model

import "time"

type (
	MstFile struct {
		ID         int64     `json:"id,omitempty"`
		ObjectName string    `json:"object_name,omitempty"`
		IsDeleted  int       `json:"is_deleted,omitempty"`
		UploadType int64     `json:"upload_type,omitempty"`
		CreatedAt  time.Time `json:"created_at,omitempty"`
		ModiedfAt  time.Time `json:"modified_at,omitempty"`
	}
	MstUploadType struct {
		ID           int64     `json:"id,omitempty"`
		UploadType   string    `json:"upload_type,omitempty"`
		ObjectPrefix string    `json:"object_prefix,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		ModiedfAt    time.Time `json:"modified_at,omitempty"`
	}
	// result for getuploadedFiles
	ResultUploadedFiles struct {
		ID           int64     `json:"id,omitempty"`
		ObjectName   string    `json:"object_name,omitempty"`
		IsDeleted    int       `json:"is_deleted,omitempty"`
		UploadType   string    `json:"upload_type,omitempty"`
		ObjectPrefix string    `json:"object_prefix,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		ModiedfAt    time.Time `json:"modified_at,omitempty"`
	}
	RequestGetUploadedFiles struct {
		Page        int
		ItemPerPage int
	}
)
