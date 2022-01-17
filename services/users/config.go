package users

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/redis"
)

var DefaultServiceConfig = ServiceConfig{
	Environment:    "dev",
	Port:           "5001",
	JWTSecret:      "SuperSecretJWT",
	SendgridAPIKey: "apikey",
	Postgres:       postgres.DefaultPostgresConfig,
	Redis:          redis.DefaultRedisConfig,
	Logger:         logger.DefaultLoggerConfig,
	Jaeger:         logger.DefaultJaegerConfig,
	WebappBaseURL:  "http://localhost:3000",
}

type ServiceConfig struct {
	Environment    string
	Port           string
	JWTSecret      string `mapstructure:"jwtSecret"`
	SendgridAPIKey string `mapstructure:"sendgridAPIKey" envconfig:"sendgrid_api_key"`
	Postgres       postgres.PostgresConfig
	Redis          redis.RedisConfig
	Jaeger         logger.JaegerConfig
	Logger         logger.LoggerConfig
	WebappBaseURL  string `mapstructure:"webAppBaseURL" envconfig:"webapp_base_url"`
}
