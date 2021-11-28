package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

// InitConnectionWithBackoff open grpc connection with <url> with backoff environment
func InitConnectionWithBackoff(url string, bc backoff.Config, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	crt := grpc.ConnectParams{Backoff: bc}
	opts = append(opts, grpc.WithConnectParams(crt))
	authConn, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, err
	}
	return authConn, nil
}
