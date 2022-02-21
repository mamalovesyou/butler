package victorinox

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
)

var DefaultVictorinoxConfig = VictorioxConfig{
	Environment: "dev",
	Services: ServicesConfig{
		Users:   postgres.DefaultConfig,
		Octopus: postgres.DefaultConfig,
	},
	Logger: logger.DefaultLoggerConfig,
}

type ServicesConfig struct {
	Users   postgres.Config `envPrefix:"USERS_POSTGRES_"`
	Octopus postgres.Config `envPrefix:"OCTOPUS_POSTGRES_"`
}

type VictorioxConfig struct {
	Environment string `env:"ENVIRONMENT"`
	Services    ServicesConfig
	Logger      logger.LoggerConfig `envPrefix:"LOGGER_"`
}
