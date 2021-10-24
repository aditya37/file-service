package service

import (
	"context"
	"errors"
)

func (s *service) GetFiles(ctx context.Context) (GetFilesResponse, error) {
	res, err := s.db.GetUploadedFiles(ctx)
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
