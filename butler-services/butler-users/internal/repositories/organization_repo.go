package repositories

import (
	models2 "github.com/butlerhq/butler/butler-services/butler-users/internal/models"
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
func (repo *OrganizationRepo) CreateOne(orga *models2.Organization) (*models2.Organization, error) {
	if err := repo.db.Create(orga).Error; err != nil {
		return &models2.Organization{}, err
	}
	return orga, nil
}

// FindByID an Organization in database and eager load Worspaces and Members
func (repo *OrganizationRepo) FindByID(organizationID string) (*models2.Organization, error) {
	org := &models2.Organization{}
	if err := repo.db.Model(org).Preload(clause.Associations).Where("id = ?", uuid.MustParse(organizationID)).Take(org).Error; err != nil {
		return &models2.Organization{}, err
	}
	return org, nil
}

// FindByUserID an Organization in database and eager load Worspaces and Members for a given user
func (repo *OrganizationRepo) FindByUserID(userID string) (*models2.Organization, error) {
	org := &models2.Organization{}

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		member := &models2.OrganizationMember{}
		where := &models2.OrganizationMember{
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

func (repo *OrganizationRepo) ListByUserID(userID string) ([]*models2.Organization, error) {
	result := []*models2.Organization{}
	// .Where("id IN (SELECT organization_id FROM organization_members WHERE user_id = ?)", userID)
	if err := repo.db.Preload("Workspaces").Preload("UserMembers", "user_id = ?", userID).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

// AddOrganizationMember add a OrganizationMember to an Organization
func (repo *OrganizationRepo) AddOrganizationMember(organizationID string, member *models2.OrganizationMember) (*models2.Workspace, error) {
	wk := &models2.Workspace{}
	tx := repo.db.Model(wk).Preload(clause.Associations).Where("id = ?", uuid.MustParse(organizationID)).Take(wk)
	if err := tx.Association("UserMembers").Append(member); err != nil {
		return &models2.Workspace{}, err
	}
	return wk, nil
}

// GetOrganizationMember for a given organizationID, userID pair
func (repo *OrganizationRepo) GetOrganizationMember(organizationID string, userID string) (*models2.OrganizationMember, error) {
	m := &models2.OrganizationMember{}
	where := &models2.OrganizationMember{
		OrganizationID: uuid.MustParse(organizationID),
		UserID:         uuid.MustParse(userID),
	}
	if err := repo.db.Where(where).First(&m).Error; err != nil {
		return &models2.OrganizationMember{}, err
	}
	return m, nil
}
