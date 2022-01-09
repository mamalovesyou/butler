package gateway

import (
	"github.com/butlerhq/butler/internal/logger"
)

var DefaultGatewayConfig = ServiceConfig{
	Environment:        "dev",
	Port:               "5001",
	Logger:             logger.DefaultLoggerConfig,
	Jaeger:             logger.DefaultJaegerConfig,
	UsersServiceAddr:   "users:3001",
	DashboardOriginUrl: "app.heybutler.local",
}

type ServiceConfig struct {
	Environment        string
	Port               string
	Jaeger             logger.JaegerConfig
	Logger             logger.LoggerConfig
	UsersServiceAddr   string
	DashboardOriginUrl string
}
