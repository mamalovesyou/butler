package internal

import (
	"context"
	"github.com/butlerhq/butler/butler-core/logger"
	"github.com/butlerhq/butler/butler-proto/gen/connectors"
	"github.com/butlerhq/butler/butler-proto/gen/workspace"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type WorkspaceRESTService struct {
	grpcAddr string
	grpcOpts []grpc.DialOption
}

func NewWorkspaceGatewayService(addr string, opts []grpc.DialOption) *WorkspaceRESTService {
	return &WorkspaceRESTService{
		grpcAddr: addr,
		grpcOpts: opts,
	}
}

// RegisterREST WorkspaceService to the specified mux
func (s *WorkspaceRESTService) RegisterREST(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	ctx := context.Background()
	if err := connectors.RegisterConnectorsServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		logger.Fatal(ctx, "Unable to register connectors endpoints", zap.Error(err))
		return err
	}
	if err := workspace.RegisterWorkspaceServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		logger.Fatal(ctx, "Unable to register workspace endpoints", zap.Error(err))
		return err
	}
	return nil
}

// ServiceName WorkspaceService to the specified mux
func (s *WorkspaceRESTService) ServiceName() string {
	return "services.Users"
}

// GRPCEndpoint return Workspace service grpc endpojt
func (s *WorkspaceRESTService) GRPCAddr() string {
	return s.grpcAddr
}

func (s *WorkspaceRESTService) GRPCDialOpts() []grpc.DialOption {
	return s.grpcOpts
}
