package workspace

import (
	"github.com/butlerhq/butler/butler-proto/gen/workspace"
	repositories2 "github.com/butlerhq/butler/butler-services/butler-users/internal/repositories"
	"gorm.io/gorm"
)

type WorkspaceService struct {
	OrganizationRepo *repositories2.OrganizationRepo
	WorkspaceRepo *repositories2.WorkspaceRepo
	workspace.UnimplementedWorkspaceServiceServer
}

func NewWorkspaceService(db *gorm.DB) *WorkspaceService {
	return &WorkspaceService{
		OrganizationRepo: repositories2.NewOrganizationRepo(db),
		WorkspaceRepo:    repositories2.NewWorkspaceRepo(db),
	}
}
