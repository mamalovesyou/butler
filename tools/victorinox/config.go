package victorinox

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
)

var DefaultVictorinoxConfig = VictorioxConfig{
	Environment: "dev",
	Services: ServicesPostgresConfig{
		Users:      postgres.DefaultPostgresConfig,
		Connectors: postgres.DefaultPostgresConfig,
	},
	Logger: logger.DefaultLoggerConfig,
}

type ServicesPostgresConfig struct {
	Users      postgres.PostgresConfig
	Connectors postgres.PostgresConfig
}

type VictorioxConfig struct {
	Environment string
	Services    ServicesPostgresConfig
	Logger      logger.LoggerConfig
}
