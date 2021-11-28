package registry

import (
	"embed"

	"github.com/matthieuberger/butler/internal/services"
	"github.com/matthieuberger/butler/internal/services/auth"
	"github.com/matthieuberger/butler/internal/services/workspace"
)

var MigrationsRegistry = map[string]embed.FS{
	services.AuthServiceName:      auth.EmbedMigrations,
	services.WorkspaceServiceName: workspace.EmbedMigrations,
}
