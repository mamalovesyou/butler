package cmd

import (
	"context"
	"time"

	"github.com/butlerhq/butler/internal/airbyte/destinations"

	"github.com/butlerhq/butler/services/users"

	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/redis"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/protocol/grpc"

	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Auth service",
		Long:  `Start Auth service`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			// Update logger with config and init tracer
			logger.UpdateAppLoggerWithConfig(usersConfig.Environment, &usersConfig.Logger)
			tracer, closer, err := logger.NewJaegerTracer(&usersConfig.Jaeger)
			if err != nil {
				logger.Fatalf(ctx, "Cannot create jaeger tracer: %+v", err)
			}
			logger.Info(ctx, "Jaeger connected")
			opentracing.SetGlobalTracer(tracer)
			defer closer.Close()

			// Initialize DB connection
			timeout := 5 * time.Second
			pgGorm := postgres.NewPostgresGorm(&usersConfig.Postgres)
			if err := pgGorm.ConnectLoop(timeout); err != nil {
				logger.Fatal(ctx, "Cannot connect to postgres.", zap.Error(err))
			}

			// Initialize redis
			rds := redis.NewRedisClient(&usersConfig.Redis)
			rds.CheckConnection(5)

			// Catalog
			abCfg := &usersConfig.Airbyte
			s3AirbyteConfig := destinations.NewS3DestinationConfig(
				abCfg.DestinationBucketName,
				abCfg.AWSRegion,
				abCfg.AWSS3Endpoint,
				abCfg.AWSAccessKeyID,
				abCfg.AWSAccessKeySecret)
			s3Destination := destinations.S3Destination{
				BaseConfig: s3AirbyteConfig,
			}
			catalog := destinations.NewDestinationCatalog(abCfg.AirbyteServerURL, &s3Destination)
			if err := catalog.Init(); err != nil {
				logger.Fatal(context.Background(), "Unable to initialize destinations catalog", zap.Error(err))
			}

			// Serve
			grpcServer := grpc.NewGRPCServer(usersConfig.Port, tracer)

			usersService := users.NewUsersService(&usersConfig, pgGorm.DB, rds.Client, catalog)
			usersService.RegisterGRPCServer(grpcServer.Server)

			healthService := users.NewHealthService(pgGorm.DB)
			healthService.RegisterGRPCServer(grpcServer.Server)
			grpcServer.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
