package permissions

import (
	"github.com/butlerhq/butler/internal/jwt"
	"github.com/butlerhq/butler/services/users/repositories"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type PermissionsUsecase struct {
	JWTManager *jwt.JWTManager
	UserRepo   *repositories.UserRepo
	TokenRepo  *repositories.TokenRepo
}

func NewPermissionsUsecase(db *gorm.DB, rdb *redis.Client, manager *jwt.JWTManager) *PermissionsUsecase {
	return &PermissionsUsecase{
		UserRepo:   repositories.NewUserRepo(db),
		TokenRepo:  repositories.NewRefreshTokenRepo(rdb),
		JWTManager: manager,
	}
}
