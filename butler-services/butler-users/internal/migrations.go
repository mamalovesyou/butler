package internal

import "embed"

//go:embed migrations
var EmbedMigrations embed.FS
