package gateway

import (
	"context"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/butlerhq/butler/api/services/users/v1"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/protocol/grpc/middlewares"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	ServiceName = "butler-gateway"
)

type RESTGatewayService struct {
	Config          *ServiceConfig
	Mux             *runtime.ServeMux
	GRPCDialOptions []grpc.DialOption
}

func NewRESTGatewayService(cfg *ServiceConfig, tracer opentracing.Tracer) *RESTGatewayService {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(middlewares.OpenTracingUnaryClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(middlewares.OpenTracingStreamClientInterceptor(tracer)),
		grpc.WithInsecure(),
	}

	return &RESTGatewayService{
		Config:          cfg,
		Mux:             mux,
		GRPCDialOptions: opts,
	}
}

func (gw *RESTGatewayService) RegisterGRPCServices() error {
	ctx := context.Background()

	// Register users service
	err := users.RegisterUsersServiceHandlerFromEndpoint(ctx, gw.Mux, gw.Config.UsersServiceAddr, gw.GRPCDialOptions)
	if err != nil {
		logger.Fatal(ctx, "Unable to register users service", zap.Error(err))
		return err
	}

	// Register octopus service
	err = octopus.RegisterOctopusServiceHandlerFromEndpoint(ctx, gw.Mux, gw.Config.OctopusServiceAddr, gw.GRPCDialOptions)
	if err != nil {
		logger.Fatal(ctx, "Unable to register octopus service", zap.Error(err))
		return err
	}

	return nil
}
