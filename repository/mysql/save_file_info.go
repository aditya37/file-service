package mysql

import (
	"context"
	"errors"
	"time"

	"github.com/aditya37/file-service/model"
)

func (m *mysql) SaveFileInfo(ctx context.Context, data model.MstFile) (int64, error) {
	times := time.Now()
	args := []interface{}{
		data.ObjectName,
		data.IsDeleted,
		data.UploadType,
		times,
	}
	row, err := m.db.Exec(mysqlInsertFile, args...)
	if err != nil {
		return 0, err
	}
	if isAffacted, _ := row.RowsAffected(); isAffacted == 0 {
		return 0, errors.New("Failed to save file")
	}
	id, _ := row.LastInsertId()
	return id, nil
}
