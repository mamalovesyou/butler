package users

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/butlerhq/butler/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

// HealthService has router and db instances
type HealthService struct {
	db *gorm.DB
}

func NewHealthService(db *gorm.DB) *HealthService {
	return &HealthService{db}
}

// Check does the health check and changes the status of the server based on wether the db is ready or not.
func (svc *HealthService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	// TODO: Perform sql.Ping() and return NOT_SERVING if databse not ready
	logger.Info(ctx, "health check", zap.Any("status", grpc_health_v1.HealthCheckResponse_SERVING))
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil

	//if _, err := svc.db.DB(); err != nil {
	//	logger.Info(ctx, "health check", zap.Any("status", grpc_health_v1.HealthCheckResponse_SERVING))
	//	return &grpc_health_v1.HealthCheckResponse{
	//		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	//	}, nil
	//} else {
	//	logger.Info(ctx, "Health check", zap.Any("status", grpc_health_v1.HealthCheckResponse_NOT_SERVING))
	//	return &grpc_health_v1.HealthCheckResponse{
	//		Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
	//	}, nil
	//}
}

// Watch is used by clients to receive updates when the service status changes.
// Watch only dummy implemented just to satisfy the interface.
func (svc *HealthService) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watching is not supported")
}

// RegisterGRPCServer Service to the specified grpc server
func (svc *HealthService) RegisterGRPCServer(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, svc)
}
