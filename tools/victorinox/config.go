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
	Users   postgres.Config
	Octopus postgres.Config
}

type VictorioxConfig struct {
	Environment string
	Services    ServicesConfig
	Logger      logger.LoggerConfig
}
