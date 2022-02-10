-- +goose Up
-- +goose StatementBegin
ALTER TABLE connectors ADD COLUMN IF NOT EXISTS config jsonb NOT NULL DEFAULT '{}'::jsonb;

ALTER TABLE connectors ADD COLUMN IF NOT EXISTS airbyte_workspace_id VARCHAR(255) NOT NULL;
ALTER TABLE connectors ADD COLUMN IF NOT EXISTS airbyte_source_definition_id VARCHAR(255) NOT NULL;
ALTER TABLE connectors ADD COLUMN IF NOT EXISTS airbyte_source_id VARCHAR(255) DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE connectors DROP COLUMN IF EXISTS airbyte_source_id;
ALTER TABLE connectors DROP COLUMN IF EXISTS airbyte_source_definition_id;
ALTER TABLE connectors DROP COLUMN IF EXISTS airbyte_workspace_id;
ALTER TABLE connectors DROP COLUMN IF EXISTS config;
-- +goose StatementEnd
