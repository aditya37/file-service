package mysql

const (
	mysqlInsertFile = `INSERT INTO mst_file(object_name,is_deleted,upload_type,created_at) VALUES(?,?,?,?)`
)
