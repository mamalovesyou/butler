package users

import (
	"context"
	butlerctx "github.com/butlerhq/butler/butler-core/context"
	"github.com/butlerhq/butler/butler-core/logger"
	"github.com/butlerhq/butler/proto/gen/auth"
	"github.com/butlerhq/butler/proto/gen/connectors"
	"github.com/butlerhq/butler/proto/gen/workspace"
	workspace2 "github.com/butlerhq/butler/services/users/internal/workspace/services/workspace"
	"github.com/butlerhq/butler/services/users/services"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type UseCase interface {
	RegisterUseCaseEndpoints(server *grpc.Server)
}

// Service has router and db instances
type Service struct {
	WorkspaceUseCase *workspace2.WorkspaceUseCase
	ConnectorUseCase *services.ConnectorsUseCase
	AuthClient       auth.AuthServiceClient
	workspace.UnimplementedWorkspaceServiceServer
}

// NewWorkspaceService initialize with predefined configuration
func NewWorkspaceService(config *ServiceConfig, db *gorm.DB) *Service {
	return &Service{
		WorkspaceUseCase: workspace2.NewWorkspaceUseCase(db),
		ConnectorUseCase: services.NewConnectorsUseCase(db),
		AuthClient:       connectAuth(config.AuthServiceAddr),
	}
}

func connectAuth(addr string) auth.AuthServiceClient {
	if conn, err := grpc.Dial(addr, []grpc.DialOption{
		grpc.WithInsecure(),
	}...); err != nil {
		logger.Error(context.Background(), "Unable to dial auth service", zap.Error(err))
		return nil
	} else {
		return auth.NewAuthServiceClient(conn)
	}
}

// RegisterGRPC Service to the specified grpc server
func (svc *Service) RegisterGRPC(server *grpc.Server) {
	connectors.RegisterConnectorsServiceServer(server, svc.ConnectorUseCase)
	workspace.RegisterWorkspaceServiceServer(server, svc.WorkspaceUseCase)
}

func (svc *Service) AuthFuncOverride(ctx context.Context, fullmethodName string) (context.Context, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.AuthFuncOverride")
	defer span.Finish()

	logger.Debug(ctx, "Checking auth privilleges")

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	if resp, err := svc.AuthClient.IsValidAccessToken(ctx, &auth.IsValidAccessTokenRequest{AccessToken: token}); err != nil {
		logger.Error(ctx, "Failed to validate access token", zap.Error(err))
		return ctx, ErrInvalidGrant
	} else {
		butlerctx.SetCtxTagUserID(ctx, resp.UserID)
		//logger.Debug(ctx, "Got userID ", zap.String("userId", resp.UserID))
		return ctx, nil
	}
}
