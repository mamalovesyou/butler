package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/butlerhq/butler/internal/jwt"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type TokenRepo struct {
	rdb *redis.Client
}

/* Note: tokens are stored using a key prefixed by the userID. This is intented so that
   we can fetch all access or refresh tokens for a given user if needed
*/

// NewUserRepo create a new repo for models.User
func NewRefreshTokenRepo(r *redis.Client) *TokenRepo {
	return &TokenRepo{rdb: r}
}

const (
	accessTokenPrefixKey  = "access-tokens"
	refreshTokenPrefixKey = "refresh-tokens"
)

func getUserAccessTokenKeyPrefix(userID string) string {
	return fmt.Sprintf("%s:%s", accessTokenPrefixKey, userID)
}
func getUserRefreshTokenKeyPrefix(userID string) string {
	return fmt.Sprintf("%s:%s", refreshTokenPrefixKey, userID)
}

// getAccessTokenKey return refresh token key for a given Subject and token unique ID
func getAccessTokenKey(userID, tokenUID string) string {
	return fmt.Sprintf("%s:%s", getUserAccessTokenKeyPrefix(userID), tokenUID)
}

// GetRefreshTokenKey return refresh token key for a given Subject and token unique ID
func getRefreshTokenKey(userID, tokenUID string) string {
	return fmt.Sprintf("%s:%s", getUserRefreshTokenKeyPrefix(userID), tokenUID)
}

// Return access token store-legacy in redis
func (repo *TokenRepo) GetAccessToken(ctx context.Context, claims *jwt.AccessTokenClaims) (string, error) {
	key := getAccessTokenKey(claims.Subject, claims.TokenID)
	token, err := repo.rdb.Get(ctx, key).Result()
	if err != nil && err == redis.Nil {
		// No refresh token is stored so this access token is revoked
		logger.Error(ctx, "Unable to find access token", zap.Error(err), zap.String("userID", claims.Subject), zap.String("tokenID", claims.TokenID))
		_, err := repo.rdb.Keys(ctx, "*").Result()
		if err != nil {
			logger.Error(ctx, "Unable to find keys", zap.Error(err))
		}
		return "", err
	}
	return token, nil
}

// Store AccessToken in redis
func (repo *TokenRepo) SetAccessToken(ctx context.Context, claims *jwt.AccessTokenClaims, token string, expiry time.Duration) error {
	key := getAccessTokenKey(claims.Subject, claims.TokenID)
	if _, err := repo.rdb.SetEX(ctx, key, token, expiry).Result(); err != nil {
		logger.Error(ctx, "Failed to store-legacy access token", zap.Error(err))
		return err
	}
	return nil
}

// Return refresh token store-legacy in redis
func (repo *TokenRepo) GetRefreshToken(ctx context.Context, claims *jwt.RefreshTokenClaims) (string, error) {
	key := getRefreshTokenKey(claims.Subject, claims.TokenID)
	token, err := repo.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

// Store RefreshToken in redis
func (repo *TokenRepo) SetRefreshToken(ctx context.Context, claims *jwt.RefreshTokenClaims, token string, expiry time.Duration) error {
	logger.Debug(ctx, "Set refresh token", zap.String("token", token))
	key := getRefreshTokenKey(claims.Subject, claims.TokenID)
	if result, err := repo.rdb.SetEX(ctx, key, token, expiry).Result(); err != nil {
		logger.Error(ctx, "Failed to store-legacy refresh token", zap.Error(err))
		return err
	} else {
		logger.Debug(ctx, "Set refresh token result", zap.String("result", result))
	}
	return nil
}

// RevokeToken remove all access and refresh tokens for a given user ID
func (repo *TokenRepo) RevokeToken(ctx context.Context, userID, tokenID string) error {
	accessTokenKey := getAccessTokenKey(userID, tokenID)
	err := repo.rdb.Del(ctx, accessTokenKey).Err()
	if err != nil {
		logger.Error(ctx, "Redis: Unable to delete access token", zap.Error(err))
		return err
	}
	return nil
}
