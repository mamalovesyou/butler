package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal             = status.Errorf(codes.Internal, "Internal error")
	ErrUnknowEmail          = status.Errorf(codes.PermissionDenied, "Unknown email")
	ErrUserNotFound         = status.Errorf(codes.NotFound, "User not found")
	ErrInvalidPassword      = status.Errorf(codes.PermissionDenied, "Password is invalid")
	ErrEmailAlreadyExists   = status.Errorf(codes.InvalidArgument, "This email is already used")
	ErrExpiredAccessToken   = status.Errorf(codes.PermissionDenied, "Invalid grant: This access token is expired")
	ErrInvalidAccessToken   = status.Errorf(codes.PermissionDenied, "Invalid grant: This access token is invalid")
	ErrRefreshTokenNotFound = status.Errorf(codes.PermissionDenied, "Invalid grant: Refresh token is revoked")
	ErrInvalidRefreshToken  = status.Errorf(codes.PermissionDenied, "Invalid grant: This refresh token is invalid")
	ErrExpiredRefreshToken  = status.Errorf(codes.PermissionDenied, "This refresh token has been revoked")
)
