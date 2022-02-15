-- +goose Up
-- +goose StatementBegin
ALTER TABLE connectors ADD COLUMN IF NOT EXISTS is_synced BOOLEAN DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE connectors DROP COLUMN IF EXISTS is_synced;
-- +goose StatementEnd
