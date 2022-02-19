package user

import (
	"github.com/butlerhq/butler/internal/jwt"
	"github.com/butlerhq/butler/services/users/repositories"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserUsecase struct {
	JWTManager       *jwt.JWTManager
	UserRepo         *repositories.UserRepo
	WorkspaceRepo    *repositories.WorkspaceRepo
	OrganizationRepo *repositories.OrganizationRepo
	InvitationRepo   *repositories.InvitationRepo
	TokenRepo        *repositories.TokenRepo
}

func NewUserUsecase(db *gorm.DB, rdb *redis.Client, manager *jwt.JWTManager) *UserUsecase {
	return &UserUsecase{
		UserRepo:       repositories.NewUserRepo(db),
		InvitationRepo: repositories.NewInvitationRepo(db),
		TokenRepo:      repositories.NewRefreshTokenRepo(rdb),
		JWTManager:     manager,
	}
}
