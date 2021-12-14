package users

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/redis"
)

type ServiceConfig struct {
	Environment     string
	Port            string
	JWTSecret       string `mapstructure:"jwtSecret"`
	Postgres        *postgres.PostgresConfig
	Redis           *redis.RedisConfig
	Jaeger          *logger.JaegerConfig
	Logger          *logger.LoggerConfig
	AuthServiceAddr string
}
