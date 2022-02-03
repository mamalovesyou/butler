package users

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/redis"
)

var DefaultServiceConfig = ServiceConfig{
	Environment:      "dev",
	Port:             "5001",
	JWTSecret:        "SuperSecretJWT",
	SendgridAPIKey:   "supersecretkey",
	Postgres:         postgres.DefaultConfig,
	Redis:            redis.DefaultRedisConfig,
	Logger:           logger.DefaultLoggerConfig,
	Jaeger:           logger.DefaultJaegerConfig,
	WebappBaseURL:    "http://localhost:3000",
	AirbyteServerURL: "airbyte-server:8081",
}

type ServiceConfig struct {
	Environment      string
	Port             string
	JWTSecret        string `mapstructure:"jwtSecret"`
	SendgridAPIKey   string `mapstructure:"sendgridAPIKey" env:"SENDGRID_API_KEY"`
	Postgres         postgres.Config
	Redis            redis.RedisConfig
	Jaeger           logger.JaegerConfig
	Logger           logger.LoggerConfig
	WebappBaseURL    string `mapstructure:"webAppBaseURL" env:"WEBAPP_BASE_URL"`
	AirbyteServerURL string `mapstructure:"airbyteServerURL" env:"AIRBYTE_SERVER_URL"`
}
