package grpc

import (
	"context"
	"fmt"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/protocol/grpc/middlewares"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// GRPCService is an interface used to register grpc services
type GRPCService interface {
	RegisterGRPC(*grpc.Server)
}

type Interceptors struct {
	Unary  []grpc.UnaryServerInterceptor
	Stream []grpc.StreamServerInterceptor
}

type GRPCServer struct {
	Port     string
	Server   *grpc.Server
	Services []GRPCService
}

func NewGRPCServer(port string, services []GRPCService, tracer opentracing.Tracer) *GRPCServer {

	// Make sure that log statements services to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLogger(logger.GetLogger())

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			middlewares.RecoveryStream(),
			middlewares.CtxTagStreamServerInterceptor(),
			middlewares.OpenTracingStreamServerInterceptor(tracer),
			middlewares.LoggerStreamServerInterceptor(logger.GetLogger()),
			grpc_validator.StreamServerInterceptor(),
			middlewares.AuthStream(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middlewares.RecoveryUnary(),
			middlewares.CtxTagUnaryServerInterceptor(),
			middlewares.OpenTracingUnaryServerInterceptor(tracer),
			middlewares.LoggerUnaryServerInterceptor(logger.GetLogger()),
			grpc_validator.UnaryServerInterceptor(),
			middlewares.AuthUnary(),
		)),
	}

	return &GRPCServer{
		Port:     port,
		Server:   grpc.NewServer(opts...),
		Services: services,
	}
}

func (srv *GRPCServer) Serve() {
	ctx := context.Background()
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	port := srv.Port
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		logger.Fatalf(ctx, "Fail to listen on port %s. %+v", port, err)
	}

	for _, service := range srv.Services {
		service.RegisterGRPC(srv.Server)
	}

	// Adding prometheus
	grpc_prometheus.Register(srv.Server)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go srv.gracefulShutdown(quit, done)

	logger.Infof(ctx, "Starting server on port %s", port)
	if err := srv.Server.Serve(listen); err != nil {
		logger.Fatalf(ctx, "Fail to serve on port %s. %+v", port, err)
	}
}

func (srv *GRPCServer) gracefulShutdown(quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Info(context.Background(), "Shutting down")
	srv.Server.GracefulStop()
	close(done)
}
