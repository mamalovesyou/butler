package cmd

import (
	"context"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/protocol/rest"
	"github.com/butlerhq/butler/internal/services/auth"
	"github.com/butlerhq/butler/internal/services/gateway"
	"github.com/butlerhq/butler/internal/services/workspace"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
			gatewayService := gateway.NewRESTAPIGatewayService(cfg)
			restSrvCfg := &rest.RESTServerConfig{
				Port:   cfg.Port,
				Mux:    gatewayService.Mux,
				Tracer: tracer,
				// TODO: R
				AllowedOrigins: []string{cfg.DashboardOriginUrl},
			}

			server := rest.NewRESTServer(restSrvCfg)

			// Register services endpoints
			// TODO: Add credentials for grpc communication
			opts := append(server.GRPCClientOpts, grpc.WithInsecure())
			authGwService := auth.NewAuthGatewayService(cfg.AuthServiceAddr, opts)
			workspaceGwService := workspace.NewWorkspaceGatewayService(cfg.WorkspaceServiceAddr, opts)
			if err := gatewayService.RegisterGRPCEndpoints(authGwService, workspaceGwService); err != nil {
				logger.Fatal(ctx, "Failed to register services", zap.Error(err))
			}

			server.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
