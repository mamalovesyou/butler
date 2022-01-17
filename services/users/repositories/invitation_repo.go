package repositories

import (
	"time"

	"github.com/google/uuid"

	"github.com/butlerhq/butler/services/users/models"
	"gorm.io/gorm"
)

type InvitationRepo struct {
	db *gorm.DB
}

// NewInvitationRepo create a new repo for models.User
func NewInvitationRepo(db *gorm.DB) *InvitationRepo {
	return &InvitationRepo{db: db}
}

func (repo *InvitationRepo) WithTransaction(f func(repo *InvitationRepo) error) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		return f(NewInvitationRepo(tx))
	})
}

// CreateOrganizationInvitations and save it in database
func (repo *InvitationRepo) CreateOrganizationInvitations(organizationID string, emails []string, expiresAt time.Time) ([]models.OrganizationInvitation, error) {
	orgaUUID := uuid.MustParse(organizationID)
	invites := make([]models.OrganizationInvitation, len(emails))
	for i, email := range emails {
		invites[i] = models.OrganizationInvitation{OrganizationID: orgaUUID, Email: email, ExpiresAt: expiresAt}
	}
	if err := repo.db.Create(&invites).Error; err != nil {
		return []models.OrganizationInvitation{}, err
	}
	return invites, nil
}

// CreateWorkspaceInvitations and save it in database
func (repo *InvitationRepo) CreateWorkspaceInvitations(workspaceID string, emails []string, expiresAt time.Time) ([]models.WorkspaceInvitation, error) {
	workspaceUUID := uuid.MustParse(workspaceID)
	invites := make([]models.WorkspaceInvitation, len(emails))
	for i, email := range emails {
		invites[i] = models.WorkspaceInvitation{WorkspaceID: workspaceUUID, Email: email, ExpiresAt: expiresAt}
	}
	if err := repo.db.Create(&invites).Error; err != nil {
		return []models.WorkspaceInvitation{}, err
	}
	return invites, nil
}
