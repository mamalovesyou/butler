package usecases

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal = status.Errorf(codes.Internal, "InternalError")
	ErrInvalidConnectorName = status.Errorf(codes.InvalidArgument, "Invalid connector name")
	ErrUnableToExchangeCode = status.Errorf(codes.Unauthenticated, "Unable to exchange code")
)
