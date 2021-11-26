package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aditya37/file-service/model"
)

func (m *mysql) GetDetailFile(ctx context.Context, object string) (*model.MstFile, error) {
	args := []interface{}{
		object,
	}

	row := m.db.QueryRowContext(ctx, "SELECT id,created_at,object_name FROM mst_file WHERE object_name = ? AND is_deleted = 0", args...)
	var record model.MstFile
	if err := row.Scan(
		&record.ID,
		&record.CreatedAt,
		&record.ObjectName,
	); err != nil {
		if err == sql.ErrNoRows {
			return &model.MstFile{}, errors.New("Data not found")
		}
		return &model.MstFile{}, err
	}

	return &record, nil
}
