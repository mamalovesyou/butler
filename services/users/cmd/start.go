package cmd

import (
	"context"
	"github.com/butlerhq/butler/services/users"
	"time"

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

			// Load config
			cfg, err := users.LoadConfig(configDir, configFileName)
			if err != nil {
				logger.Fatalf(ctx, "Failed to load config: %+v", err)
			}
			logger.Debug(ctx, "Starting Users service", zap.Any("config", cfg))

			// Update logger with config and init tracer
			logger.UpdateAppLoggerWithConfig(cfg.Logger)
			tracer, closer, err := logger.NewJaegerTracer(cfg.Jaeger)
			if err != nil {
				logger.Fatalf(ctx, "Cannot create jaeger tracer: %+v", err)
			}
			logger.Info(ctx, "Jaeger connected")
			opentracing.SetGlobalTracer(tracer)
			defer closer.Close()

			// Initialize DB connection
			timeout := 5 * time.Second
			pgGorm := postgres.NewPostgresGorm(cfg.Postgres)
			if err := pgGorm.ConnectLoop(timeout); err != nil {
				logger.Fatal(ctx, "Cannot connect to postgres.", zap.Error(err))
			}

			// Initialize redis
			rdb := redis.NewRedisClient(cfg.Redis)

			// Serve
			grpcServer := grpc.NewGRPCServer(cfg.Port, tracer)
			usersService := users.NewUsersService(cfg, pgGorm.DB, rdb)
			usersService.RegisterGRPCServer(grpcServer.Server)
			grpcServer.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
