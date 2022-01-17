package repositories

import (
	"time"

	"gorm.io/gorm/clause"

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

func (repo *InvitationRepo) DB() *gorm.DB {
	return repo.db
}

func (repo *InvitationRepo) WithTransaction(db *gorm.DB) *InvitationRepo {
	return NewInvitationRepo(db)
}

// GetInvitation retrieves an invitation by ID and token
func (repo *InvitationRepo) GetInvitation(invitationID, token string) (models.Invitation, error) {
	invitation := models.Invitation{Token: token, BaseModel: models.BaseModel{ID: uuid.MustParse(invitationID)}}
	if err := repo.db.Preload(clause.Associations).First(&invitation).Error; err != nil {
		return models.Invitation{}, err
	}
	return invitation, nil
}

// CreateOrganizationInvitations will create an invitation for a given organization for each emails
func (repo *InvitationRepo) CreateOrganizationInvitations(organizationID string, emails []string, expiresAt time.Time) ([]models.Invitation, error) {
	orgaUUID := uuid.MustParse(organizationID)
	invites := make([]models.Invitation, len(emails))
	for i, email := range emails {
		invites[i] = models.Invitation{OrganizationID: orgaUUID, Email: email, ExpiresAt: expiresAt}
	}
	if err := repo.db.Create(&invites).Error; err != nil {
		return []models.Invitation{}, err
	}
	return invites, nil
}

// CreateWorkspaceInvitations will create an invitation for a given organization for each emails
func (repo *InvitationRepo) CreateWorkspaceInvitations(organizationID string, workspaceID string, emails []string, expiresAt time.Time) ([]models.Invitation, error) {
	invites := make([]models.Invitation, len(emails))
	for i, email := range emails {
		invites[i] = models.Invitation{
			OrganizationID: uuid.MustParse(organizationID),
			WorkspaceID:    uuid.MustParse(workspaceID),
			Email:          email,
			ExpiresAt:      expiresAt,
		}
	}
	if err := repo.db.Create(&invites).Error; err != nil {
		return []models.Invitation{}, err
	}
	return invites, nil
}

// DeleteInvitation will soft delete an invitation
func (repo *InvitationRepo) DeleteInvitation(invitationID string) error {
	invitation := models.Invitation{BaseModel: models.BaseModel{ID: uuid.MustParse(invitationID)}}
	if err := repo.db.Delete(&invitation).Error; err != nil {
		return err
	}
	return nil
}
