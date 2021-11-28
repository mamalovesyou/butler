package middlewares

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

// AddValidator returns grpc.Server config option that turn on payload validation.
func AddValidator(opts []grpc.ServerOption) []grpc.ServerOption {

	// Add unary interceptor
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_validator.UnaryServerInterceptor(),
	))

	// Add stream interceptor (added as an example here)
	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_validator.StreamServerInterceptor(),
	))

	return opts
}
