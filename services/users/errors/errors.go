package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknownEmail         = status.Errorf(codes.InvalidArgument, "No users registered with this address")
	ErrInvalidPassword      = status.Errorf(codes.InvalidArgument, "Email/Password combination")
	ErrEmailAlreadyUsed     = status.Errorf(codes.InvalidArgument, "This email is already used")
	ErrMissingAccessToken   = status.Errorf(codes.Unauthenticated, "Invalid grant: Missing access token")
	ErrExpiredAccessToken   = status.Errorf(codes.Unauthenticated, "Invalid grant: This access token is expired")
	ErrInvalidAccessToken   = status.Errorf(codes.Unauthenticated, "Invalid grant: This access token is invalid")
	ErrRefreshTokenNotFound = status.Errorf(codes.Unauthenticated, "Invalid grant: Refresh token is revoked")
	ErrInvalidRefreshToken  = status.Errorf(codes.Unauthenticated, "Invalid grant: This refresh token is invalid")
	ErrExpiredRefreshToken  = status.Errorf(codes.Unauthenticated, "This refresh token has been revoked")
	ErrUserNotFound         = status.Errorf(codes.InvalidArgument, "This user dosen't exists")
)
