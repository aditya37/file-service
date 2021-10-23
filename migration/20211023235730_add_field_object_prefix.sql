-- +goose Up
-- +goose StatementBegin
ALTER TABLE mst_upload_type ADD COLUMN object_prefix VARCHAR(32) NOT NULL AFTER id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_upload_type DROP COLUMN object_prefix;
-- +goose StatementEnd
