package victorinox

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/butlerhq/butler/services/octopus"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/services/users"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

var (
	UsersMigrationName   = "users"
	OctopusMigrationName = "octopus"

	AllowedCommands = []string{
		"up", "up-by-one", "up-to", "down", "down-to", "redo", "reset", "status", "version", "fix",
	}
)

type PostgresMigrationsPair struct {
	Migreations    embed.FS
	PostgresConfig *postgres.Config
}

type GooseMigrations map[string]*PostgresMigrationsPair

func NewGooseMigrations(config *VictorioxConfig) *GooseMigrations {
	migrations := make(GooseMigrations)

	// Adding user service migration
	migrations[UsersMigrationName] = &PostgresMigrationsPair{
		Migreations:    users.EmbedMigrations,
		PostgresConfig: &config.Services.Users,
	}

	// Adding octopus service migration
	migrations[OctopusMigrationName] = &PostgresMigrationsPair{
		Migreations:    octopus.EmbedMigrations,
		PostgresConfig: &config.Services.Octopus,
	}

	return &migrations
}

func IsSupportedGooseCmd(cmd string) bool {
	sort.Strings(AllowedCommands)
	i := sort.SearchStrings(AllowedCommands, cmd)
	return i < len(AllowedCommands) && AllowedCommands[i] == cmd
}

func (m *GooseMigrations) RunGooseMigrationForAllServices(ctx context.Context, cmd string, args ...string) error {
	for service, _ := range *m {
		if err := m.RunGooseMigrationForService(ctx, service, cmd, args...); err != nil {
			logger.Fatal(ctx, "Failed to run migrations", zap.Error(err), zap.String("service", service))
		}
	}
	return nil
}

func (m *GooseMigrations) RunGooseMigrationForService(ctx context.Context, service, cmd string, args ...string) error {
	service = strings.TrimSpace(service)
	config, exists := (*m)[service]
	if !exists {
		logger.Fatal(ctx, "Unknown service", zap.String("service", service))
		return errors.New(fmt.Sprintf("Unknown service: %s", service))
	}

	logger.Infof(ctx, "Applying migrations %s", service)
	// Initialize DB connection
	pg := postgres.NewPostgresGorm(config.PostgresConfig)
	if err := pg.ConnectLoop(5 * time.Second); err != nil {
		logger.Fatal(ctx, "Cannot connect to postgres.", zap.Error(err), zap.String("service", service))
	}

	goose.SetBaseFS(config.Migreations)
	if err := goose.Run(cmd, pg.SqlDB, "migrations", args...); err != nil {
		logger.Fatalf(ctx, "Goose %v: %v", cmd, err)
	}

	return nil
}
