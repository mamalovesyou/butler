package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type RESTAPIGatewayService struct {
	Config      *ServiceConfig
	Mux         *runtime.ServeMux
	DialOptions []grpc.DialOption
}

type RESTService interface {
	ServiceName() string
	RegisterREST(*runtime.ServeMux, string, []grpc.DialOption) error
	GRPCAddr() string
	GRPCDialOpts() []grpc.DialOption
}

func NewRESTAPIGatewayService(cfg *ServiceConfig) *RESTAPIGatewayService {
	mux := runtime.NewServeMux()
	return &RESTAPIGatewayService{
		Config:      cfg,
		Mux:         mux,
		DialOptions: []grpc.DialOption{grpc.WithInsecure()},
	}
}

func (gw *RESTAPIGatewayService) RegisterGRPCServices() error {
	ctx := context.Background()
	if err := RegisterUsersService(ctx, gw.Mux, gw.Config.AuthServiceAddr, gw.DialOptions); err != nil {
		return err
	}
	return nil
}
