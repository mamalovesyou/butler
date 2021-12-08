package workspace

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/redis"
	"github.com/butlerhq/butler/internal/services"
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
	if err := viper.UnmarshalKey(services.WorkspaceServiceName, svcConfig); err != nil {
		return &ServiceConfig{}, err
	}
	return svcConfig, nil
}
