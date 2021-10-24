package mysql

import (
	"context"

	"github.com/aditya37/file-service/model"
)

func (m *mysql) GetUploadedFiles(ctx context.Context) ([]*model.ResultUploadedFiles, error) {
	rows, err := m.db.QueryContext(ctx, mysqlGetUploadedFile)
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
