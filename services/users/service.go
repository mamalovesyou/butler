package users

import (
	api_users "github.com/butlerhq/butler/api/services/users/v1"
	"github.com/butlerhq/butler/internal/jwt"
	"github.com/butlerhq/butler/services/users/usecase/permissions"
	"github.com/butlerhq/butler/services/users/usecase/user"
	"github.com/butlerhq/butler/services/users/usecase/workspace"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// UsersService has router and db instances
type UsersService struct {
	UserUsecase       *user.UserUsecase
	WorkspaceUsecase  *workspace.WorkspaceUsecase
	PermissionUseCase *permissions.PermissionsUsecase

	api_users.UnimplementedUsersServiceServer
}

// NewUsersService initialize with predefined configuration
func NewUsersService(config *ServiceConfig, db *gorm.DB, rdb *redis.Client) *UsersService {
	jwtManager := jwt.NewJWTManager(config.JWTSecret)
	return &UsersService{
		WorkspaceUsecase:  workspace.NewWorkspaceUsecase(db, config.SendgridAPIKey, config.WebappBaseURL),
		UserUsecase:       user.NewUserUsecase(db, rdb, jwtManager),
		PermissionUseCase: permissions.NewPermissionsUsecase(db, rdb, jwtManager),
	}
}

// RegisterGRPCServer Service to the specified grpc server
func (svc *UsersService) RegisterGRPCServer(server *grpc.Server) {
	api_users.RegisterUsersServiceServer(server, svc)
}
