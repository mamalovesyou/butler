package middlewares

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// Return the opentracing unary server interceptor
func OpenTracingUnaryServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	o := []grpc_opentracing.Option{
		grpc_opentracing.WithTracer(tracer),
	}

	return grpc_middleware.ChainUnaryServer(grpc_opentracing.UnaryServerInterceptor(o...))
}

// Return the opentracing stream server interceptor
func OpenTracingStreamServerInterceptor(tracer opentracing.Tracer) grpc.StreamServerInterceptor {
	o := []grpc_opentracing.Option{
		grpc_opentracing.WithTracer(tracer),
	}

	return grpc_middleware.ChainStreamServer(grpc_opentracing.StreamServerInterceptor(o...))
}

// Return the opentracing unary client interceptor
func OpenTracingUnaryClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	o := []grpc_opentracing.Option{
		grpc_opentracing.WithTracer(tracer),
	}

	return grpc_middleware.ChainUnaryClient(grpc_opentracing.UnaryClientInterceptor(o...))
}

// Return the opentracing stream client interceptor
func OpenTracingStreamClientInterceptor(tracer opentracing.Tracer) grpc.StreamClientInterceptor {
	o := []grpc_opentracing.Option{
		grpc_opentracing.WithTracer(tracer),
	}

	return grpc_middleware.ChainStreamClient(grpc_opentracing.StreamClientInterceptor(o...))
}
