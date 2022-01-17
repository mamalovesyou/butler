package victorinox

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
)

var DefaultVictorinoxConfig = VictorioxConfig{
	Environment: "dev",
	Services: ServicesPostgresConfig{
		Users:   postgres.DefaultPostgresConfig,
		Octopus: postgres.DefaultPostgresConfig,
	},
	Logger: logger.DefaultLoggerConfig,
}

type ServicesPostgresConfig struct {
	Users   postgres.PostgresConfig
	Octopus postgres.PostgresConfig
}

type VictorioxConfig struct {
	Environment string
	Services    ServicesPostgresConfig
	Logger      logger.LoggerConfig
}
