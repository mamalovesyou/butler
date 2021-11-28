package grpc

import (
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// initialize client with tracing interceptor using grpc client side chaining
func NewClientConn(addr string, tracer opentracing.Tracer) (*grpc.ClientConn, error) {

	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithCodes(codes.Unavailable, codes.Aborted),
		grpc_retry.WithMax(10),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
		grpc_retry.WithPerRetryTimeout(time.Second * 5),
	}

	return grpc.Dial(
		addr,
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_retry.StreamClientInterceptor(retryOpts...),
			grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_retry.UnaryClientInterceptor(retryOpts...),
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
		grpc.WithInsecure(),
	)
}
