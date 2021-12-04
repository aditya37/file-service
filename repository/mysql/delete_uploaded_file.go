package mysql

import (
	"context"
	"errors"
)

const mysqlQueryDeleteFile = `UPDATE mst_file mf SET mf.is_deleted = 1 WHERE mf.object_name = ?`

func (m *mysql) DeleteUploadedFile(ctx context.Context, obj string) error {
	row, err := m.db.ExecContext(ctx, mysqlQueryDeleteFile, obj)
	if err != nil {
		return err
	}
	if isAffacted, _ := row.RowsAffected(); isAffacted == 0 {
		return errors.New("Failed delete file")
	}
	return nil
}
