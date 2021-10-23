package mysql

import (
	"context"

	"github.com/aditya37/file-service/model"
)

func (m *mysql) GetUploadType(ctx context.Context) ([]*model.MstUploadType, error) {

	rows, err := m.db.Query("SELECT id,object_prefix,upload_type FROM mst_upload_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.MstUploadType
	for rows.Next() {
		var record model.MstUploadType
		if err := rows.Scan(
			&record.ID,
			&record.ObjectPrefix,
			&record.UploadType,
		); err != nil {
			return nil, err
		}
		result = append(result, &model.MstUploadType{
			ID:           record.ID,
			ObjectPrefix: record.ObjectPrefix,
			UploadType:   record.UploadType,
		})
	}
	return result, nil
}
