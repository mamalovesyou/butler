-- +goose Up
-- +goose StatementBegin
ALTER TABLE workspaces ADD COLUMN IF NOT EXISTS airbyte_workspace_id VARCHAR(255) DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE workspaces DROP COLUMN IF EXISTS airbyte_workspace_id;
-- +goose StatementEnd