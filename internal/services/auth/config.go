package auth

import (
	"github.com/matthieuberger/butler/internal/logger"
	"github.com/matthieuberger/butler/internal/postgres"
	"github.com/matthieuberger/butler/internal/redis"
	"github.com/matthieuberger/butler/internal/services"
	"github.com/matthieuberger/butler/internal/utils"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Environment string
	Port        string
	Postgres    *postgres.PostgresConfig
	JWTSecret   string `mapstructure:"jwtSecret"`
	Redis       *redis.RedisConfig
	Jaeger      *logger.JaegerConfig
	Logger      *logger.LoggerConfig
}

// Load and return a *auth.ServiceConfig if no error
func LoadConfig(path, name string) (*ServiceConfig, error) {
	if err := utils.LoadConfigWithViper(path, name); err != nil {
		return &ServiceConfig{}, err
	}
	svcConfig := &ServiceConfig{}
	if err := viper.UnmarshalKey(services.AuthServiceName, svcConfig); err != nil {
		return &ServiceConfig{}, err
	}
	return svcConfig, nil
}
