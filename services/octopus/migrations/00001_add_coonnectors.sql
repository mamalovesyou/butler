-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'auth_scheme_enum') THEN
CREATE TYPE auth_scheme_enum AS ENUM (
            'OAUTH2',
            'API_KEY'
        );
END IF;
END
$$;

CREATE TABLE IF NOT EXISTS connectors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workspace_id UUID NOT NULL,
    provider VARCHAR(255) NOT NULL,
    auth_scheme auth_scheme_enum NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
                             );

CREATE INDEX IF NOT EXISTS connectors_idx ON connectors(workspace_id);
CREATE UNIQUE INDEX connectors_provider_idx ON connectors(workspace_id, provider);

CREATE TABLE IF NOT EXISTS connector_secrets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    connector_id UUID NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
                             );

CREATE UNIQUE INDEX IF NOT EXISTS connector_secrets_connector_idx ON connector_secrets(connector_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS connectors_provider_idx;
DROP INDEX IF EXISTS connectors_idx;
DROP INDEX IF EXISTS connector_secrets_connector_idx;

DROP TABLE IF EXISTS connectors_secrets;
DROP TABLE IF EXISTS onnectors;

DROP TYPE auth_scheme_enum;
-- +goose StatementEnd