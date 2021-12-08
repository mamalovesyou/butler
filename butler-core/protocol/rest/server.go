package rest

import (
	"context"
	"fmt"
	"github.com/butlerhq/butler/butler-core/protocol/grpc/middlewares"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/butlerhq/butler/butler-core/logger"
	"github.com/butlerhq/butler/butler-core/protocol/rest/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

type RESTServerConfig struct {
	Port           string
	Mux            *runtime.ServeMux
	Tracer         opentracing.Tracer
	AllowedOrigins []string
}

type RESTServer struct {
	server         *http.Server
	port           string
	GRPCClientOpts []grpc.DialOption
}

type GRPCClientInterceptors struct {
	Unary  grpc.UnaryClientInterceptor
	Stream grpc.StreamClientInterceptor
}

func NewRESTServer(cfg *RESTServerConfig) *RESTServer {

	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(middlewares.OpenTracingUnaryClientInterceptor(cfg.Tracer)),
		grpc.WithStreamInterceptor(middlewares.OpenTracingStreamClientInterceptor(cfg.Tracer)),
	}

	c := cors.New(cors.Options{
		AllowedOrigins: cfg.AllowedOrigins,
		AllowedHeaders: []string{
			"App-Version",
			"Content-Type",
			"Origin",
			"Accept",
			"Authorization",
		},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	corsHandler := c.Handler(cfg.Mux)
	endpointsHandler := middleware.AddOpenTracing(middleware.AddLogger(logger.GetLogger(), corsHandler))

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", endpointsHandler)

	return &RESTServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.Port),
			Handler: mux,
		},
		port:           cfg.Port,
		GRPCClientOpts: opts,
	}

}

func (srv *RESTServer) gracefullShutdown(quit <-chan os.Signal, done chan<- bool) {
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Info(ctx, "Server is shutting down...")

	if err := srv.server.Shutdown(ctx); err != nil {
		logger.Info(ctx, "Could not gracefully shutdown the server...", zap.Error(err))
	}
	close(done)
}

func (srv *RESTServer) Serve() {

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	ctx := context.Background()

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go srv.gracefullShutdown(quit, done)
	logger.Info(ctx, "Starting HTTP/REST API Gateway...", zap.String("port", srv.port))

	if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(ctx, "Failed to listen", zap.String("port", srv.port), zap.Error(err))
	}

	<-done
	logger.Info(ctx, "Server stopped")
}
