package workspace

import (
	"github.com/butlerhq/butler/services/users/repositories"
	"gorm.io/gorm"
)

type WorkspaceUsecase struct {
	OrganizationRepo *repositories.OrganizationRepo
	WorkspaceRepo    *repositories.WorkspaceRepo
}

func NewWorkspaceUsecase(db *gorm.DB) *WorkspaceUsecase {
	return &WorkspaceUsecase{
		OrganizationRepo: repositories.NewOrganizationRepo(db),
		WorkspaceRepo:    repositories.NewWorkspaceRepo(db),
	}
}
