package gateway

import (
	"github.com/butlerhq/butler/internal/logger"
)

var DefaultGatewayConfig = ServiceConfig{
	Environment:        "dev",
	Port:               "5001",
	Logger:             logger.DefaultLoggerConfig,
	Jaeger:             logger.DefaultJaegerConfig,
	UsersServiceAddr:   "users:3001",
	OctopusServiceAddr: "octopus:3002",
	WebAppOriginUrl:    "*",
}

type ServiceConfig struct {
	Environment        string              `env:"ENVIRONMENT"`
	Port               string              `env:"PORT"`
	Jaeger             logger.JaegerConfig `envPrefix:"JAEGER_"`
	Logger             logger.LoggerConfig `envPrefix:"LOGGER_"`
	UsersServiceAddr   string              `env:"USERS_SERVICE_ADDR"`
	OctopusServiceAddr string              `env:"OCTOPUS_SERVICE_ADDR"`
	WebAppOriginUrl    string              `env:"WEBAPP_ORIGIN_URL"`
}
