package auth

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	butlerctx "github.com/butlerhq/butler/internal/context"
	"github.com/butlerhq/butler/internal/jwt"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/services/auth/models"
	"github.com/butlerhq/butler/internal/services/auth/repositories"
	"github.com/butlerhq/butler/internal/services/gen/auth"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// Service has router and db instances
type Service struct {
	JWTManager *jwt.JWTManager
	UserRepo   *repositories.UserRepo
	TokenRepo  *repositories.TokenRepo
	auth.UnimplementedAuthServiceServer
}

// NewAuthService initialize with predefined configuration
func NewAuthService(config *ServiceConfig, db *gorm.DB, rdb *redis.Client) *Service {
	return &Service{
		UserRepo:   repositories.NewUserRepo(db),
		TokenRepo:  repositories.NewRefreshTokenRepo(rdb),
		JWTManager: jwt.NewJWTManager(config.JWTSecret),
	}
}

// RegisterGRPC Service to the specified grpc server
func (s *Service) RegisterGRPC(server *grpc.Server) {
	auth.RegisterAuthServiceServer(server, s)
}

// auth create access and refresh tokens for a given user
func (s *Service) auth(ctx context.Context, user *models.User) (*auth.AuthenticatedUser, error) {
	// Create token pair
	sub := user.ID.String()
	accessTokanClaims := s.JWTManager.NewAccessTokenClaims(sub, user.Email, user.FirstName, user.LastName)
	accessToken, err := s.JWTManager.CreateAccessToken(accessTokanClaims)
	if err != nil {
		logger.Error(ctx, "Failed to create access token", zap.Error(err))
		return &auth.AuthenticatedUser{}, ErrInternal
	}
	refreshTokenClaims := s.JWTManager.NewRefreshTokenClaims(sub)
	refreshToken, err := s.JWTManager.CreateRefreshToken(refreshTokenClaims)
	if err != nil {
		logger.Error(ctx, "Failed to create refresh token", zap.Error(err))
		return &auth.AuthenticatedUser{}, ErrInternal
	}

	// Store accessToken and refresh token
	if err := s.TokenRepo.SetAccessToken(ctx, accessTokanClaims, accessToken, s.JWTManager.AccessTokenDuration); err != nil {
		logger.Error(ctx, "Failed to store-legacy refresh token", zap.Error(err))
		return &auth.AuthenticatedUser{}, ErrInternal
	}
	if err := s.TokenRepo.SetRefreshToken(ctx, refreshTokenClaims, refreshToken, s.JWTManager.RefreshTokenDuration); err != nil {
		logger.Error(ctx, "Failed to store-legacy refresh token", zap.Error(err))
		return &auth.AuthenticatedUser{}, ErrInternal
	}

	return &auth.AuthenticatedUser{
		User:         user.ToPb(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// IsValidAccessToken return IsValidAccessTokenResponse if given access token is valid
func (s *Service) IsValidAccessToken(ctx context.Context, req *auth.IsValidAccessTokenRequest) (*auth.IsValidAccessTokenResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auth.IsValidAccessToken")
	defer span.Finish()
	// Verify if access token is valid
	claims, err := s.JWTManager.IsValidAccessToken(req.AccessToken)
	if err != nil {
		return &auth.IsValidAccessTokenResponse{}, ErrExpiredAccessToken
	}
	butlerctx.SetCtxTagUserID(ctx, claims.Subject)

	// Verify if token is not revoked and match record
	token, err := s.TokenRepo.GetAccessToken(ctx, claims)
	if err != nil {
		return &auth.IsValidAccessTokenResponse{}, ErrInternal
	}
	if token != req.AccessToken {
		return &auth.IsValidAccessTokenResponse{}, ErrInvalidAccessToken
	}

	return &auth.IsValidAccessTokenResponse{UserID: claims.Subject}, nil
}

// SignIn auth a user with an email/passsword combination
func (s *Service) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auth.SignIn")
	defer span.Finish()
	user, err := s.UserRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &auth.AuthenticatedUser{}, ErrUnknowEmail
		} else {
			return &auth.AuthenticatedUser{}, ErrInternal
		}
	}
	if err = s.UserRepo.VerifyPassword(user, req.Password); err != nil {
		return &auth.AuthenticatedUser{}, ErrInvalidPassword
	}

	return s.auth(ctx, user)
}

// SignUp creates a new user
func (s *Service) SignUp(ctx context.Context, req *auth.SignUpRequest) (resp *auth.AuthenticatedUser, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auth.SignUp")
	defer span.Finish()
	user := &models.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		HashPassword: s.UserRepo.HashPassword(req.Password),
	}

	if user, err = s.UserRepo.CreateOne(user); err != nil {
		if postgres.IsDuplicateKeyError(err) {
			return &auth.AuthenticatedUser{}, ErrEmailAlreadyExists
		} else {
			return &auth.AuthenticatedUser{}, ErrInternal
		}
	}

	return s.auth(ctx, user)
}

//// SignOut revoke a user tokens
//func (s *Service) SignOut(ctx context.Context, req *auth.SignOutRequest) (*emptypb.Empty, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "auth.SignOut")
//	defer span.Finish()
//
//	claims, err := s.JWTManager.IsValidAccessToken(req.AccessToken)
//	if err != nil {
//		logger.Error(ctx, "Invalid access token", zap.Error(err))
//		return &emptypb.Empty{}, ErrInvalidAccessToken
//	}
//
//	if err := s.TokenRepo.RevokeToken(ctx, claims.Subject, claims.TokenID); err != nil {
//		logger.Error(ctx, "Unable to revoke access token", zap.Error(err))
//		return &emptypb.Empty{}, ErrInternal
//	}
//
//	return &emptypb.Empty{}, nil
//}

// RefreshToken
func (s *Service) RefreshToken(ctx context.Context, req *auth.RefreshRequest) (*auth.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auth.RefreshToken")
	defer span.Finish()

	// Verify if refresh token is valid
	claims, err := s.JWTManager.IsValidRefreshToken(req.RefreshToken)
	if err != nil {
		logger.Errorf(ctx, "Failed to validate refresh token. %+v", err)
		return &auth.AuthenticatedUser{}, ErrExpiredRefreshToken
	}

	// Verify if token is not revoked.
	logger.Debugf(ctx, "Comparing with request token", zap.String("req token", req.RefreshToken))
	token, err := s.TokenRepo.GetRefreshToken(ctx, claims)
	logger.Debugf(ctx, "Found existing refresh token", zap.String("token", token), zap.Error(err))
	if err == redis.Nil {
		return &auth.AuthenticatedUser{}, ErrRefreshTokenNotFound
	} else if token != req.RefreshToken {
		return &auth.AuthenticatedUser{}, ErrExpiredRefreshToken
	} else if err != nil {
		return &auth.AuthenticatedUser{}, ErrInternal
	}

	user, err := s.UserRepo.FindByID(claims.Subject)
	if err != nil {
		logger.Errorf(ctx, "Unable to fetch user. %+v", err)
		return &auth.AuthenticatedUser{}, ErrUserNotFound
	}

	return s.auth(ctx, user)
}

// ListUsers return a list of users for a given userID list
func (s *Service) ListUsers(ctx context.Context, req *auth.ListUsersRequest) (*auth.ListUsersResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auth.ListUsers")
	defer span.Finish()

	users, err := s.UserRepo.ListByIDs(req.UserIDs)
	if err != nil {
		return &auth.ListUsersResponse{}, ErrInternal
	}

	result := make([]*auth.User, len(users))
	for i, u := range users {
		result[i] = u.ToPb()
	}

	return &auth.ListUsersResponse{
		Users: result,
	}, nil
}
