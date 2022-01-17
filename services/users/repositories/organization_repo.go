package repositories

import (
	"github.com/butlerhq/butler/services/users/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationRepo struct {
	db *gorm.DB
}

// NewOrganizationRepo create a new repo for models.User
func NewOrganizationRepo(db *gorm.DB) *OrganizationRepo {
	return &OrganizationRepo{db: db}
}

func (repo *OrganizationRepo) WithTransaction(db *gorm.DB) *OrganizationRepo {
	return NewOrganizationRepo(db)
}

// CreateOne and save it in database
func (repo *OrganizationRepo) CreateOne(orga *models.Organization) (*models.Organization, error) {
	if err := repo.db.Create(orga).Error; err != nil {
		return &models.Organization{}, err
	}
	return orga, nil
}

// UpdateOne and save updates it in database
func (repo *OrganizationRepo) UpdateOne(organizationID string, updates models.Organization) (*models.Organization, error) {
	orga := &models.Organization{}
	if err := repo.db.Model(orga).Where("id = ?", organizationID).Updates(updates).Error; err != nil {
		return &models.Organization{}, err
	}
	return orga, nil
}

// FindByID an Organization in database and eager load Worspaces and Members
func (repo *OrganizationRepo) FindByID(organizationID string) (*models.Organization, error) {
	org := &models.Organization{}
	if err := repo.db.Model(org).Preload(clause.Associations).Where("id = ?", uuid.MustParse(organizationID)).Take(org).Error; err != nil {
		return &models.Organization{}, err
	}
	return org, nil
}

// FindByUserID an Organization in database and eager load Worspaces and Members for a given user
func (repo *OrganizationRepo) FindByUserID(userID string) (*models.Organization, error) {
	org := &models.Organization{}

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		member := &models.OrganizationMember{}
		where := &models.OrganizationMember{
			UserID: uuid.MustParse(userID),
		}
		if err := repo.db.Where(where).First(&member).Error; err != nil {
			return err
		}
		if err := repo.db.Model(org).Preload(clause.Associations).Where("id = ?", member.OrganizationID).Take(org).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return org, nil
}

func (repo *OrganizationRepo) ListByUserID(userID string) ([]models.Organization, error) {
	result := []models.Organization{}
	// .Where("id IN (SELECT organization_id FROM organization_members WHERE user_id = ?)", userID)
	if err := repo.db.Preload("Workspaces").Preload("UserMembers", "user_id = ?", userID).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

// AddOrganizationMember add a OrganizationMember to an Organization
func (repo *OrganizationRepo) AddOrganizationMember(organizationID string, member *models.OrganizationMember) (*models.Workspace, error) {
	wk := &models.Workspace{}
	tx := repo.db.Model(wk).Preload(clause.Associations).Where("id = ?", uuid.MustParse(organizationID)).Take(wk)
	if err := tx.Association("UserMembers").Append(member); err != nil {
		return &models.Workspace{}, err
	}
	return wk, nil
}

// GetOrganizationMember for a given organizationID, userID pair
func (repo *OrganizationRepo) GetOrganizationMember(organizationID string, userID string) (*models.OrganizationMember, error) {
	m := &models.OrganizationMember{}
	where := &models.OrganizationMember{
		OrganizationID: uuid.MustParse(organizationID),
		UserID:         uuid.MustParse(userID),
	}
	if err := repo.db.Where(where).First(&m).Error; err != nil {
		return &models.OrganizationMember{}, err
	}
	return m, nil
}
