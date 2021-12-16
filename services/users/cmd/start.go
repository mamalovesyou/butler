package cmd

import (
	"context"
	"time"

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
			logger.UpdateAppLoggerWithConfig(&usersConfig.Logger)
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
			rdb := redis.NewRedisClient(&usersConfig.Redis)

			// Serve
			grpcServer := grpc.NewGRPCServer(usersConfig.Port, tracer)
			usersService := users.NewUsersService(&usersConfig, pgGorm.DB, rdb)
			usersService.RegisterGRPCServer(grpcServer.Server)
			grpcServer.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
