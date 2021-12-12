package user

import (
	"context"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/errors"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// RefreshToken
func (uc *UserUsecase) RefreshToken(ctx context.Context, refreshToken string) (*models.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.RefreshToken")
	defer span.Finish()

	// Verify if refresh token is valid
	claims, err := uc.JWTManager.IsValidRefreshToken(refreshToken)
	if err != nil {
		logger.Errorf(ctx, "Failed to validate refresh token. %+v", err)
		return nil, errors.ErrExpiredRefreshToken
	}

	// Verify if token is not revoked.
	token, err := uc.TokenRepo.GetRefreshToken(ctx, claims)
	if err == redis.Nil {
		return nil, errors.ErrRefreshTokenNotFound
	} else if token != refreshToken {
		return nil, errors.ErrExpiredRefreshToken
	} else if err != nil {
		logger.Error(ctx, "Unable to retrieve refresh token", zap.Error(err))
		return nil, err
	}

	user, err := uc.UserRepo.FindByID(claims.Subject)
	if err != nil {
		logger.Errorf(ctx, "Unable to fetch user. %+v", err)
		return nil, errors.ErrUserNotFound
	}

	return uc.authenticate(ctx, user)
}
