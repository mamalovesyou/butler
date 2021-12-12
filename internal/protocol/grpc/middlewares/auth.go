package middlewares

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

// EmptyAuthFunc does not do anything
// Use AuthFuncOverride if you want to overide
func emptyAuthFunc(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// AuthUnary returns grpc.UnaryServerInterceptor that turn on user.
func AuthUnary() grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(emptyAuthFunc)
}

// AuthStream returns grpc.StreamServerInterceptor that turn on user.
func AuthStream() grpc.StreamServerInterceptor {
	return grpc_auth.StreamServerInterceptor(emptyAuthFunc)
}
