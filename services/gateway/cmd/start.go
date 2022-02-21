package cmd

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/protocol/rest"
	"github.com/butlerhq/butler/services/gateway"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Gateway server.",
		Long:  `Start Gateway server.`,
		Run: func(cmd *cobra.Command, args []string) {

			ctx := context.Background()

			// Update logger with config and init tracer
			logger.UpdateAppLoggerWithConfig(gatewayCfg.Environment, &gatewayCfg.Logger)
			tracer, closer, err := logger.NewJaegerTracer(&gatewayCfg.Jaeger)
			if err != nil {
				logger.Fatalf(ctx, "Cannot create jaeger tracer: %+v", err)
			}
			logger.Info(ctx, "Jaeger connected")
			opentracing.SetGlobalTracer(tracer)
			defer closer.Close()

			// Create gateway instance
			gatewayService := gateway.NewRESTGatewayService(&gatewayCfg, tracer)
			gatewayService.RegisterGRPCServices()

			// Create rest http server
			serverCfg := &rest.RESTServerConfig{
				Port:           gatewayCfg.Port,
				Mux:            gatewayService.Mux,
				AllowedOrigins: []string{gatewayCfg.WebAppOriginUrl},
			}

			server := rest.NewRESTServer(serverCfg)
			server.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
