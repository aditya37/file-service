-- +goose Up
-- +goose StatementBegin
CREATE TABLE mst_upload_type (
    id BIGINT(10) NOT NULL AUTO_INCREMENT,
    upload_type VARCHAR(32) NOT NULL DEFAULT "PHOTO_PROFILE",
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NULL,
    UNIQUE(upload_type),
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE mst_upload_type;
-- +goose StatementEnd
