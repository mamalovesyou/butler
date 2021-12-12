DROP EXTENSION IF EXISTS "uuid-ossp";

DROP TABLE IF EXISTS organizations;

DROP TABLE IF EXISTS workspaces;
DROP INDEX IF EXISTS workspaces_organization_idx;

DROP TABLE IF EXISTS organization_members;
DROP INDEX IF EXISTS organization_members_organization_idx;

DROP TABLE IF EXISTS workspace_members;
DROP INDEX IF EXISTS workspace_members_workspace_idx;