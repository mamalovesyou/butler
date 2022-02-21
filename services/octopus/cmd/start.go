package cmd

import (
	"context"
	"time"

	"github.com/butlerhq/butler/internal/airbyte/sources/catalog"

	"github.com/butlerhq/butler/services/octopus/services"

	"github.com/butlerhq/butler/services/octopus"

	"github.com/butlerhq/butler/internal/postgres"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/protocol/grpc"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Octopus service",
		Long:  `Start Octopus service`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			// Update logger with config and init tracer
			logger.UpdateAppLoggerWithConfig(cfgService.Environment, &cfgService.Logger)
			tracer, closer, err := logger.NewJaegerTracer(&cfgService.Jaeger)
			if err != nil {
				logger.Fatalf(ctx, "Cannot create jaeger tracer: %+v", err)
			}
			logger.Info(ctx, "Jaeger connected")
			opentracing.SetGlobalTracer(tracer)
			defer closer.Close()

			// Initialize DB connection
			timeout := 5 * time.Second
			pgGorm := postgres.NewPostgresGorm(&cfgService.Postgres)
			if err := pgGorm.ConnectLoop(timeout); err != nil {
				logger.Fatal(ctx, "Cannot connect to postgres.", zap.Error(err))
			}

			// Catalog
			catalog := catalog.NewCatalog(&cfgService.Sources, cfgService.AirbyteServerURL)
			if err := catalog.Init(); err != nil {
				logger.Fatal(context.Background(), "Unable to initialize catalog", zap.Error(err))
			}

			// Serve
			grpcServer := grpc.NewGRPCServer(cfgService.Port, tracer)

			dataSourceSvc := services.NewDataSourcesService(catalog)
			dataSourceSvc.RegisterGRPCServer(grpcServer.Server)

			airbyteSvc := services.NewAirbyteService(cfgService.AirbyteServerURL)
			connectorSvc := services.NewConnectorsService(catalog, pgGorm.DB, airbyteSvc)
			connectorSvc.RegisterGRPCServer(grpcServer.Server)

			octopusService := octopus.NewOctopusService(pgGorm.DB, catalog)
			octopusService.RegisterGRPCServer(grpcServer.Server)

			healthService := octopus.NewHealthService(pgGorm.DB)
			healthService.RegisterGRPCServer(grpcServer.Server)

			grpcServer.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
