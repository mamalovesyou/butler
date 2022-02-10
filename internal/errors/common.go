package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal         = status.Error(codes.Internal, "Internal error")
	ErrInvalidArguments = status.Error(codes.InvalidArgument, "Invalid arguments")
	ErrInvalidGrant     = status.Error(codes.Unauthenticated, "Invalid grant")
)
