package middlewares

import (
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

// AddLogging returns grpc.Server config option that turn on logging.
func CtxTagUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	o := []grpc_ctxtags.Option{
		grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
	}

	return grpc_ctxtags.UnaryServerInterceptor(o...)
}

// AddLogging returns grpc.Server config option that turn on logging.
func CtxTagStreamServerInterceptor() grpc.StreamServerInterceptor {
	o := []grpc_ctxtags.Option{
		grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
	}
	return grpc_ctxtags.StreamServerInterceptor(o...)
}
