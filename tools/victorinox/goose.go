package victorinox

import (
	"context"
	"embed"
	"sort"
	"strings"
	"time"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/services/users"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

var (
	UsersMigrationName = "users"

	AllowedCommands = []string{
		"up", "up-by-one", "up-to", "down", "down-to", "redo", "reset", "status", "version", "fix",
	}
)

type GooseMigrations struct {
	dbConfig     postgres.PostgresConfig
	migrationMap map[string]embed.FS
}

func NewGooseMigrations(config *VictorioxConfig) *GooseMigrations {
	migrationMap := make(map[string]embed.FS)
	migrationMap[UsersMigrationName] = users.EmbedMigrations
	return &GooseMigrations{
		dbConfig:     config.Postgres,
		migrationMap: migrationMap,
	}
}

func IsSupportedGooseCmd(cmd string) bool {
	sort.Strings(AllowedCommands)
	i := sort.SearchStrings(AllowedCommands, cmd)
	return i < len(AllowedCommands) && AllowedCommands[i] == cmd
}

func (m *GooseMigrations) RunGooseMigration(ctx context.Context, name, cmd string, args ...string) error {
	name = strings.TrimSpace(name)
	logger.Infof(ctx, "Available migrations %+v", m.migrationMap)
	logger.Infof(ctx, "Applying migrations %s", name)

	// Initialize DB connection
	pg := postgres.NewPostgresGorm(&m.dbConfig)
	if err := pg.ConnectLoop(5 * time.Second); err != nil {
		logger.Fatal(ctx, "Cannot connect to postgres.", zap.Error(err))
	}

	if embedded, ok := m.migrationMap[name]; ok {
		goose.SetBaseFS(embedded)
		if err := goose.Run(cmd, pg.SqlDB, "migrations", args...); err != nil {
			logger.Fatalf(ctx, "Goose %v: %v", cmd, err)
		}
	} else {
		logger.Fatalf(ctx, "Embeded migrations %s not found", name)
	}

	return nil
}