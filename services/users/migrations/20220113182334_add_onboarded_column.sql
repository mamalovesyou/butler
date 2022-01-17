-- +goose Up
-- +goose StatementBegin
ALTER TABLE organizations ADD COLUMN IF NOT EXISTS onboarded BOOLEAN DEFAULT FALSE NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE organizations DROP COLUMN IF EXISTS onboarded;
-- +goose StatementEnd
