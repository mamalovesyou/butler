package victorinox

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
)

var DefaultVictorinoxConfig = VictorioxConfig{
	Environment: "dev",
	Postgres:    postgres.DefaultPostgresConfig,
	Logger:      logger.DefaultLoggerConfig,
}

type VictorioxConfig struct {
	Environment string
	Postgres    postgres.PostgresConfig
	Logger      logger.LoggerConfig
}
