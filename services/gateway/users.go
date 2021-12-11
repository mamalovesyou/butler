package gateway

import (
	"context"
	"github.com/butlerhq/butler/api/services/users/v1"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterUsersService(ctx context.Context, mux *runtime.ServeMux, grpcAddr string, opts []grpc.DialOption) error {
	err := users.RegisterUsersServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		logger.Fatal(ctx, "Unable to register users service", zap.Error(err))
		return err
	}
	return nil
}
