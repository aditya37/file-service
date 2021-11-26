package mysql

const (
	mysqlInsertFile      = `INSERT INTO mst_file(object_name,is_deleted,upload_type,created_at) VALUES(?,?,?,?)`
	mysqlGetUploadedFile = `SELECT * FROM (
		SELECT 
			mst_file.id,
			mst_file.object_name,
			mst_file.is_deleted,
			mst_upload_type.object_prefix,
			mst_upload_type.upload_type,
			mst_file.created_at
		FROM 
			mst_file 
		INNER JOIN mst_upload_type ON mst_file.upload_type = mst_upload_type.id WHERE mst_file.is_deleted = 0
	) AS ls ORDER BY ls.created_at DESC LIMIT ?,?`
)
