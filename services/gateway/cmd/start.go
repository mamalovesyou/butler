package cmd

import (
	"context"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/protocol/rest"
	"github.com/butlerhq/butler/services/gateway"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Gateway server.",
		Long:  `Start Gateway server.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()

			//Load config
			cfg, err := gateway.LoadConfig(configDir, configFileName)
			if err != nil {
				logger.Fatalf(ctx, "Failed to load config: %+v", err)
			}
			logger.Info(ctx, "Starting Gateway microservice", zap.Any("config", cfg))

			// Update logger with config and init tracer
			logger.UpdateAppLoggerWithConfig(cfg.Logger)
			tracer, closer, err := logger.NewJaegerTracer(cfg.Jaeger)
			if err != nil {
				logger.Fatalf(ctx, "Cannot create jaeger tracer: %+v", err)
			}
			logger.Info(ctx, "Jaeger connected")
			opentracing.SetGlobalTracer(tracer)
			defer closer.Close()

			// Create gateway instance
			gatewayService := gateway.NewRESTGatewayService(cfg, tracer)
			gatewayService.RegisterGRPCServices()

			// Create rest http server
			serverCfg := &rest.RESTServerConfig{
				Port:           cfg.Port,
				Mux:            gatewayService.Mux,
				AllowedOrigins: []string{cfg.DashboardOriginUrl},
			}

			server := rest.NewRESTServer(serverCfg)
			server.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
