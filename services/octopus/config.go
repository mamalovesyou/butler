package octopus

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/services/octopus/connectors"
)

var DefaultServiceConfig = ServiceConfig{
	Environment: "dev",
	Port:        "5002",
	Logger:      logger.DefaultLoggerConfig,
	Jaeger:      logger.DefaultJaegerConfig,
}

type ServiceConfig struct {
	Environment string
	Port        string
	Connectors  connectors.ConnectorsConfig
	Postgres    postgres.PostgresConfig
	Jaeger      logger.JaegerConfig
	Logger      logger.LoggerConfig
}
