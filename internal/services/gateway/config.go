package gateway

import (
	"github.com/matthieuberger/butler/internal/logger"
	"github.com/matthieuberger/butler/internal/services"
	"github.com/matthieuberger/butler/internal/utils"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Environment          string
	Port                 string
	Jaeger               *logger.JaegerConfig
	Logger               *logger.LoggerConfig
	AuthServiceAddr      string
	WorkspaceServiceAddr string
	DashboardOriginUrl   string
}

// Load and return a *auth.ServiceConfig if no error
func LoadConfig(path, name string) (*ServiceConfig, error) {
	if err := utils.LoadConfigWithViper(path, name); err != nil {
		return &ServiceConfig{}, err
	}
	svcConfig := &ServiceConfig{}
	if err := viper.UnmarshalKey(services.GatewayServiceName, svcConfig); err != nil {
		return &ServiceConfig{}, err
	}
	return svcConfig, nil
}
