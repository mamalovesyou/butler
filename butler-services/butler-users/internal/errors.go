package internal

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal             = status.Errorf(codes.Internal, "Internal error")
	ErrMissingIndentity     = status.Errorf(codes.PermissionDenied, "Missing identity")
	ErrInvalidGrant         = status.Errorf(codes.PermissionDenied, "Invalid grant")
	ErrOrganizationNotFound = status.Errorf(codes.NotFound, "Organization not found")
	ErrWorkspaceNotFound    = status.Errorf(codes.NotFound, "Workspace not found")
)
