package users

import (
	"github.com/butlerhq/butler/internal/airbyte"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/redis"
)

var DefaultServiceConfig = ServiceConfig{
	Environment:    "dev",
	Port:           "5001",
	JWTSecret:      "SuperSecretJWT",
	SendgridAPIKey: "supersecretkey",
	Postgres:       postgres.DefaultConfig,
	Redis:          redis.DefaultRedisConfig,
	Logger:         logger.DefaultLoggerConfig,
	Jaeger:         logger.DefaultJaegerConfig,
	WebappBaseURL:  "http://localhost:3000",
	Airbyte:        airbyte.DefaultAirbyteConfig,
}

type ServiceConfig struct {
	Environment    string              `env:"ENVIRONMENT"`
	Port           string              `env:"PORT"`
	JWTSecret      string              `mapstructure:"jwtSecret"`
	SendgridAPIKey string              `mapstructure:"sendgridAPIKey" env:"SENDGRID_API_KEY"`
	Postgres       postgres.Config     `envPrefix:"POSTGRES_"`
	Redis          redis.Config        `envPrefix:"REDIS_"`
	Jaeger         logger.JaegerConfig `envPrefix:"JAEGER_"`
	Logger         logger.LoggerConfig `envPrefix:"LOGGER_"`
	WebappBaseURL  string              `mapstructure:"webAppBaseURL" env:"WEBAPP_BASE_URL"`
	Airbyte        airbyte.Config
}
