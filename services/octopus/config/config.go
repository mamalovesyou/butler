package config

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
)

var DefaultServiceConfig = ServiceConfig{
	Environment:      "dev",
	Port:             "5002",
	Logger:           logger.DefaultLoggerConfig,
	Jaeger:           logger.DefaultJaegerConfig,
	AirbyteServerURL: "http://airbyte-server:8081",
}

type ServiceConfig struct {
	Environment      string
	Port             string
	Sources          SourcesConfig
	Postgres         postgres.Config
	Jaeger           logger.JaegerConfig
	Logger           logger.LoggerConfig
	AirbyteServerURL string `mapStructure:"redirectURL" env:"AIRBYTE_SERVER_URL"`
}
