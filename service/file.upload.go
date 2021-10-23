package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"

	"github.com/aditya37/file-service/utils"
	"github.com/google/uuid"
)

// Create tempory file Before Upload
func (f *service) storeToTemp(file multipart.File, fileName string) (string, error) {

	// save file to tempory before upload
	prefixTmpFile := fmt.Sprintf("*.%s", fileName)

	tmp, err := ioutil.TempFile(os.TempDir(), prefixTmpFile)
	if err != nil {
		return "", &utils.CustomError{
			InternalError: err.Error(),
			Description:   "Failed create temp file",
			Code:          ErrCodeFailedCreateTemp,
		}
	}
	defer tmp.Close()

	// file byte
	fb, err := ioutil.ReadAll(file)
	if err != nil {
		return "", &utils.CustomError{
			InternalError: err.Error(),
			Description:   "Failed to read file",
			Code:          ErrCodeReadFile,
		}
	}

	// write file to temp file
	if _, err := tmp.Write(fb); err != nil {
		return "", &utils.CustomError{
			InternalError: err.Error(),
			Description:   "Failed to write file to temp",
			Code:          ErrCodeFailedWriteTemp,
		}
	}

	return tmp.Name(), nil
}

// Function for create unique object name
func (f *service) generateObjectName(uploadType string) string {
	// unix object name
	uuid := uuid.New()
	timeStamp := time.Now().Unix()
	return fmt.Sprintf("%s.%s.%d", uploadType, uuid, timeStamp)
}

// firebase handling
func (f *service) firebaseHandling(ctx context.Context, file multipart.File, fileName, objectName string) (size int64, obj, medialink string, err error) {

	tmpFile, err := f.storeToTemp(file, fileName)
	if err != nil {
		return 0, "", "", err
	}

	// get object name after upload
	upload, err := f.storage.Upload(ctx, tmpFile, objectName)
	if err != nil {
		return 0, "", "", &utils.CustomError{
			InternalError: err.Error(),
			Description:   "Failed save data to server",
			Code:          ErrCodeFailedUploadToServer,
		}
	}

	// get object attributes
	attrs, err := f.storage.GetObjectAttribute(ctx, upload)
	if err != nil {
		return 0, "", "", &utils.CustomError{
			InternalError: err.Error(),
			Description:   "Failed to parse object attribute",
			Code:          ErrCodeFailedParseObject,
		}
	}

	return attrs.Size, upload, attrs.MediaLink, nil
}

// process Upload photoProfile
func (f *service) processUpload(ctx context.Context, req ProcessUpload) (FileUploadResponse, error) {
	generatedObj := f.generateObjectName(req.UploadType)

	size, obj, medialink, err := f.firebaseHandling(ctx, req.File, req.Filename, generatedObj)
	if err != nil {
		return FileUploadResponse{}, err
	}

	return FileUploadResponse{
		Id:         0,
		MediaLink:  medialink,
		ObjectName: obj,
		FileSize:   size,
	}, nil
}

// Main Businnes logic
func (f *service) FileUpload(ctx context.Context, request FileUploadRequest) (FileUploadResponse, error) {

	// validate upload type
	if request.UploadType == UploadTypeContent {
		// validate content type
		if request.FileDetail.ContentType != "image/png" && request.FileDetail.ContentType != "image/jpeg" {
			return FileUploadResponse{}, &utils.CustomError{
				InternalError: "please use file type image/png or image/jpeg",
				Description:   "Wrong file format type",
				Code:          ErrCodeWrongFileFormat,
			}
		}
		// handle upload photo profile
		resp, err := f.processUpload(ctx, ProcessUpload{
			FileType:   request.FileDetail.ContentType,
			UploadType: "photo.profile",
			Filename:   request.FileDetail.FileName,
			File:       request.FileDetail.File,
		})
		if err != nil {
			return FileUploadResponse{}, err
		}
		return resp, nil
	} else if request.UploadType == UploadTypePhotoProfile {
		// validate content type
		if request.FileDetail.ContentType != "image/png" && request.FileDetail.ContentType != "image/jpeg" {
			return FileUploadResponse{}, &utils.CustomError{
				InternalError: "please use file type image/png or image/jpeg",
				Description:   "Wrong file format type",
				Code:          ErrCodeWrongFileFormat,
			}
		}
		// handle upload photo content
		resp, err := f.processUpload(ctx, ProcessUpload{
			FileType:   request.FileDetail.ContentType,
			UploadType: "photo.content",
			Filename:   request.FileDetail.FileName,
			File:       request.FileDetail.File,
		})
		if err != nil {
			return FileUploadResponse{}, err
		}
		return resp, nil
	} else if request.UploadType == UploadTypeDocument {
		if request.FileDetail.ContentType != "application/pdf" && request.FileDetail.ContentType != "application/vnd.ms-excel" && request.FileDetail.ContentType != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" {
			return FileUploadResponse{}, &utils.CustomError{
				InternalError: "please use file type pdf,xls,docx",
				Description:   "Wrong file format type",
				Code:          ErrCodeWrongFileFormat,
			}
		}
		resp, err := f.processUpload(ctx, ProcessUpload{
			FileType:   request.FileDetail.ContentType,
			UploadType: "document",
			Filename:   request.FileDetail.FileName,
			File:       request.FileDetail.File,
		})
		if err != nil {
			return FileUploadResponse{}, err
		}
		return resp, nil
	} else {
		return FileUploadResponse{}, &utils.CustomError{
			InternalError: "Wrong request",
			Description:   "Wrong request upload type",
			Code:          ErrCodeWrongRequest,
		}
	}
}
