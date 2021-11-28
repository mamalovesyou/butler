package middlewares

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DefaultRecoveryHandler() grpc_recovery.RecoveryHandlerFunc {
	// Panic handler prints the stack trace when recovering from a panic.
	return grpc_recovery.RecoveryHandlerFunc(func(p interface{}) error {
		return status.Errorf(codes.Internal, "%s", p)
	})

}

// Recovery(Unary returns grpc.UnaryServerInterceptor that turn on Recovery(.
func RecoveryUnary() grpc.UnaryServerInterceptor {
	o := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(DefaultRecoveryHandler()),
	}

	return grpc_recovery.UnaryServerInterceptor(o...)
}

// Recovery(Stream returns grpc.StreamServerInterceptor that turn on Recovery(.
func RecoveryStream() grpc.StreamServerInterceptor {
	o := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(DefaultRecoveryHandler()),
	}

	return grpc_recovery.StreamServerInterceptor(o...)
}
