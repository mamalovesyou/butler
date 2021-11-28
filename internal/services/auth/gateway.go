package auth

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/matthieuberger/butler/internal/services"
	"github.com/matthieuberger/butler/internal/services/gen/auth"
	"google.golang.org/grpc"
)

type AuthRESTService struct {
	grpcAddr string
	grpcOpts []grpc.DialOption
}

func NewAuthGatewayService(addr string, opts []grpc.DialOption) *AuthRESTService {
	return &AuthRESTService{
		grpcAddr: addr,
		grpcOpts: opts,
	}
}

// RegisterREST AuthService to the specified mux
func (s *AuthRESTService) RegisterREST(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return auth.RegisterAuthServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

// ServiceName AuthService to the specified mux
func (s *AuthRESTService) ServiceName() string {
	return services.AuthServiceName
}

// GRPCEndpoint return auth service grpc endpojt
func (s *AuthRESTService) GRPCAddr() string {
	return s.grpcAddr
}

func (s *AuthRESTService) GRPCDialOpts() []grpc.DialOption {
	return s.grpcOpts
}
