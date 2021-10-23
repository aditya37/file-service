-- +goose Up
-- +goose StatementBegin
ALTER TABLE mst_file ADD CONSTRAINT `fk_mst_file_upload_type` FOREIGN KEY (`upload_type`) REFERENCES `mst_upload_type` (`id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_file DROP CONSTRAINT fk_mst_file_upload_type;
-- +goose StatementEnd
