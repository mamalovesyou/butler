package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal     = status.Error(codes.Internal, "Internal error")
	ErrInvalidGrant = status.Error(codes.Unauthenticated, "Invalid grant")
)
