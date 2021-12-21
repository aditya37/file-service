package service

import (
	"context"
	"errors"

	"github.com/aditya37/file-service/model"
)

func (s *service) GetFiles(ctx context.Context, request GetFileRequest) (GetFilesResponse, error) {

	// validate pagination request
	if request.Page == 0 {
		request.Page = 1
	}
	if request.ItemPerPage == 0 {
		request.ItemPerPage = 5
	}
	res, err := s.db.GetUploadedFiles(ctx, model.RequestGetUploadedFiles{
		Page:        request.Page,
		ItemPerPage: request.ItemPerPage,
	})
	if err != nil {
		return GetFilesResponse{}, err
	}
	if len(res) == 0 {
		return GetFilesResponse{}, errors.New("Empty")
	}

	var fl []*FileItems
	for _, val := range res {
		metaData, err := s.storage.GetObject(ctx, val.ObjectName)
		if err != nil {
			return GetFilesResponse{}, err
		}
		fl = append(fl, &FileItems{
			Id:           val.ID,
			CreatedAt:    val.CreatedAt,
			UploadType:   val.UploadType,
			ObjectPrefix: val.ObjectPrefix,
			IsDeleted:    val.IsDeleted,
			Metadata: Metadata{
				ContentType: metaData.ContentType,
				MediaLink:   metaData.MediaLink,
				FileSize:    metaData.Size,
			},
		})
	}

	return GetFilesResponse{
		Count:     int64(len(res)),
		FileItems: fl,
	}, nil
}
