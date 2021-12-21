-- +goose Up
-- +goose StatementBegin
CREATE TABLE mst_file (
    id BIGINT(10) NOT NULL AUTO_INCREMENT,
    object_name VARCHAR(255) NOT NULL,
    is_deleted TINYINT(1) NOT NULL DEFAULT '0',
    upload_type BIGINT(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NULL,
    UNIQUE(object_name),
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE mst_file;
-- +goose StatementEnd
