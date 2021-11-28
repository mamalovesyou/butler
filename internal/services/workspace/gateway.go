package workspace

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/matthieuberger/butler/internal/services"
	"github.com/matthieuberger/butler/internal/services/gen/workspace"
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
	return workspace.RegisterWorkspaceServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

// ServiceName WorkspaceService to the specified mux
func (s *WorkspaceRESTService) ServiceName() string {
	return services.WorkspaceServiceName
}

// GRPCEndpoint return Workspace service grpc endpojt
func (s *WorkspaceRESTService) GRPCAddr() string {
	return s.grpcAddr
}

func (s *WorkspaceRESTService) GRPCDialOpts() []grpc.DialOption {
	return s.grpcOpts
}
