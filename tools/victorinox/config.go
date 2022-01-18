package victorinox

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
)

var DefaultVictorinoxConfig = VictorioxConfig{
	Environment: "dev",
	Postgres:    postgres.DefaultConfig,
	Logger:      logger.DefaultLoggerConfig,
}

type ServicesConfig struct {
	Users   postgres.Config
	Octopus postgres.Config
}

type VictorioxConfig struct {
	Environment string
	Postgres    postgres.Config
	Logger      logger.LoggerConfig
}
