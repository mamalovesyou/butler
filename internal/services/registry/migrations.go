package registry

import (
	"embed"

	"github.com/butlerhq/butler/internal/services"
	"github.com/butlerhq/butler/internal/services/auth"
	"github.com/butlerhq/butler/internal/services/workspace"
)

var MigrationsRegistry = map[string]embed.FS{
	services.AuthServiceName:      auth.EmbedMigrations,
	services.WorkspaceServiceName: workspace.EmbedMigrations,
}
