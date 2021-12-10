package workspace

import (
	"github.com/butlerhq/butler/proto/gen/workspace"
	"github.com/butlerhq/butler/services/users/repositories"
	"gorm.io/gorm"
)

type WorkspaceService struct {
	OrganizationRepo *repositories.OrganizationRepo
	WorkspaceRepo    *repositories.WorkspaceRepo
	workspace.UnimplementedWorkspaceServiceServer
}

func NewWorkspaceService(db *gorm.DB) *WorkspaceService {
	return &WorkspaceService{
		OrganizationRepo: repositories.NewOrganizationRepo(db),
		WorkspaceRepo:    repositories.NewWorkspaceRepo(db),
	}
}
