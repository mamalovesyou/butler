package user

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"go.uber.org/zap"
)

// users create access and refresh tokens for a given user
func (uc *UserUsecase) authenticate(ctx context.Context, user *models.User) (*models.AuthenticatedUser, error) {
	// Create token pair
	sub := user.ID.String()
	accessTokenClaims := uc.JWTManager.NewAccessTokenClaims(sub, user.Email, user.FirstName, user.LastName)
	accessToken, err := uc.JWTManager.CreateAccessToken(accessTokenClaims)
	if err != nil {
		logger.Error(ctx, "Failed to create access token", zap.Error(err))
		return &models.AuthenticatedUser{}, err
	}
	refreshTokenClaims := uc.JWTManager.NewRefreshTokenClaims(sub)
	refreshToken, err := uc.JWTManager.CreateRefreshToken(refreshTokenClaims)
	if err != nil {
		logger.Error(ctx, "Failed to create refresh token", zap.Error(err))
		return &models.AuthenticatedUser{}, err
	}

	// Store accessToken and refresh token
	if err := uc.TokenRepo.SetAccessToken(ctx, accessTokenClaims, accessToken, uc.JWTManager.AccessTokenDuration); err != nil {
		logger.Error(ctx, "Failed to store-legacy refresh token", zap.Error(err))
		return &models.AuthenticatedUser{}, err
	}
	if err := uc.TokenRepo.SetRefreshToken(ctx, refreshTokenClaims, refreshToken, uc.JWTManager.RefreshTokenDuration); err != nil {
		logger.Error(ctx, "Failed to store-legacy refresh token", zap.Error(err))
		return &models.AuthenticatedUser{}, err
	}

	return &models.AuthenticatedUser{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
