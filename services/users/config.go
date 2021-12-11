package users

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/redis"
	"github.com/butlerhq/butler/internal/utils"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Environment     string
	Port            string
	Postgres        *postgres.PostgresConfig
	Redis           *redis.RedisConfig
	Jaeger          *logger.JaegerConfig
	Logger          *logger.LoggerConfig
	AuthServiceAddr string
}

// Load and return a *auth.ServiceConfig if no error
func LoadConfig(path, name string) (*ServiceConfig, error) {
	if err := utils.LoadConfigWithViper(path, name); err != nil {
		return &ServiceConfig{}, err
	}
	svcConfig := &ServiceConfig{}
	// TODO: Replace configName
	if err := viper.UnmarshalKey("service.users", svcConfig); err != nil {
		return &ServiceConfig{}, err
	}
	return svcConfig, nil
}
