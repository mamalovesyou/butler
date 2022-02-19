package permissions

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/errors"
	"github.com/butlerhq/butler/services/users/models"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (uc *PermissionsUsecase) IsAuthenticated(ctx context.Context, token string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "permissions_ucase.IsAuthenticated")
	defer span.Finish()

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		logger.Error(ctx, "Unable to retrieve users token from context", zap.Error(err))
		return nil, errors.ErrMissingAccessToken
	}

	// Verify if access token is valid
	claims, err := uc.JWTManager.IsValidAccessToken(token)
	if err != nil {
		return nil, errors.ErrExpiredAccessToken
	}

	// Verify if token is not revoked
	if _, err = uc.TokenRepo.GetAccessToken(ctx, claims); err != nil {
		logger.Error(ctx, "Unable to retrieve access token", zap.Error(err))
		return nil, err
	}

	user, err := uc.UserRepo.FindByID(claims.Subject)
	if err != nil {
		logger.Error(ctx, "Uable to retrieve user", zap.Error(err))
		return nil, err
	}
	return user, nil

}
