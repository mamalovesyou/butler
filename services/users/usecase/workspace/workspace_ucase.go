package workspace

import (
	"github.com/butlerhq/butler/services/users/repositories"
	"gorm.io/gorm"
)

type WorkspaceUsecase struct {
	OrganizationRepo *repositories.OrganizationRepo
	WorkspaceRepo    *repositories.WorkspaceRepo
	InvitationRepo   *repositories.InvitationRepo
}

func NewWorkspaceUsecase(db *gorm.DB) *WorkspaceUsecase {
	return &WorkspaceUsecase{
		OrganizationRepo: repositories.NewOrganizationRepo(db),
		WorkspaceRepo:    repositories.NewWorkspaceRepo(db),
		InvitationRepo:   repositories.NewInvitationRepo(db),
	}
}
