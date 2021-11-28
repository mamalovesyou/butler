package gateway

import (
	"context"
	"github.com/matthieuberger/butler/internal/logger"
	"go.uber.org/zap"

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

func (gw *RESTAPIGatewayService) RegisterGRPCEndpoints(services ...RESTService) error {
	ctx := context.Background()
	for _, s := range services {
		logger.Debug(ctx, "Attempt to rest service", zap.String("name", s.ServiceName()))
		if err := s.RegisterREST(gw.Mux, s.GRPCAddr(), s.GRPCDialOpts()); err != nil {
			logger.Error(ctx, "Failted to register rest service", zap.String("name", s.ServiceName()), zap.Error(err))
			return err
		}
		logger.Debug(ctx, "Successfully registered rest service", zap.String("name", s.ServiceName()))
	}
	return nil
}
