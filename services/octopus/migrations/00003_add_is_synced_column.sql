-- +goose Up
-- +goose StatementBegin
ALTER TABLE connectors ADD COLUMN IF NOT EXISTS airbyte_destination_id VARCHAR(255) DEFAULT NULL;
ALTER TABLE connectors ADD COLUMN IF NOT EXISTS airbyte_connection_id VARCHAR(255) DEFAULT NULL;
ALTER TABLE connectors ADD COLUMN IF NOT EXISTS is_active BOOLEAN DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE connectors DROP COLUMN IF EXISTS schedule;
ALTER TABLE connectors DROP COLUMN IF EXISTS airbyte_connection_id;
ALTER TABLE connectors DROP COLUMN IF EXISTS airbyte_destination_id;
-- +goose StatementEnd
