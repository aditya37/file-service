package mysql

import (
	"context"

	"github.com/aditya37/file-service/model"
)

func (m *mysql) GetUploadedFiles(ctx context.Context, data model.RequestGetUploadedFiles) ([]*model.ResultUploadedFiles, error) {
	offset := (data.Page - 1) * data.ItemPerPage
	args := []interface{}{
		offset,
		data.ItemPerPage,
	}
	rows, err := m.db.QueryContext(ctx, mysqlGetUploadedFile, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.ResultUploadedFiles
	for rows.Next() {
		var record model.ResultUploadedFiles
		if err := rows.Scan(
			&record.ID,
			&record.ObjectName,
			&record.IsDeleted,
			&record.ObjectPrefix,
			&record.UploadType,
			&record.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, &model.ResultUploadedFiles{
			ID:           record.ID,
			ObjectName:   record.ObjectName,
			IsDeleted:    record.IsDeleted,
			UploadType:   record.UploadType,
			ObjectPrefix: record.ObjectPrefix,
			CreatedAt:    record.CreatedAt,
		})
	}
	return result, nil
}
