package workspace

import (
	"context"
	"time"

	postgres2 "github.com/matthieuberger/butler/internal/postgres"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/matthieuberger/butler/internal/logger"
	"github.com/matthieuberger/butler/internal/protocol/grpc"
	"github.com/matthieuberger/butler/internal/services/workspace"
	"github.com/spf13/cobra"
)

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start workspace service",
		Long:  `Start workspace service`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			// Load config
			cfg, err := workspace.LoadConfig(configDir, configFileName)
			if err != nil {
				logger.Fatalf(ctx, "Failed to load config: %+v", err)
			}
			logger.Info(ctx, "Starting workspace microservice", zap.Any("config", cfg))

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
			postgres := postgres2.NewPostgresGorm(cfg.Postgres)
			if err := postgres.ConnectLoop(timeout); err != nil {
				logger.Fatalw(ctx, "Cannot connect to postgres.", "error", err)
			}

			// Initialize redis
			// rdb := redis.NewRedisClient(cfg.Redis)
			service := workspace.NewWorkspaceService(cfg, postgres.DB)
			server := grpc.NewGRPCServer(cfg.Port, []grpc.GRPCService{service}, tracer)
			server.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
}
