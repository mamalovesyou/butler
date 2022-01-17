package repositories

import (
	"fmt"

	"github.com/butlerhq/butler/services/users/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WorkspaceRepo struct {
	db *gorm.DB
}

// NewWorkspaceRepo create a new repo for models.User
func NewWorkspaceRepo(db *gorm.DB) *WorkspaceRepo {
	return &WorkspaceRepo{db: db}
}

func (repo *WorkspaceRepo) WithTransaction(db *gorm.DB) *WorkspaceRepo {
	return NewWorkspaceRepo(db)
}

// CreateOne and save it in database
func (repo *WorkspaceRepo) CreateOne(orga *models.Workspace) (*models.Workspace, error) {
	if err := repo.db.Create(orga).Error; err != nil {
		return &models.Workspace{}, err
	}
	if err := repo.db.Save(orga).Error; err != nil {
		return &models.Workspace{}, err
	}
	return orga, nil
}

// FindByID an Workspace in database and eager load Worspaces and Members
func (repo *WorkspaceRepo) FindByID(workspaceID string) (*models.Workspace, error) {
	ws := &models.Workspace{}
	if err := repo.db.Model(ws).Preload(clause.Associations).Where("id = ?", workspaceID).Take(ws).Error; err != nil {
		return &models.Workspace{}, err
	}
	fmt.Println()
	fmt.Printf("Workspace Inivtes: %v", ws.PendingInvitations)
	fmt.Println()
	return ws, nil
}

// AddWorkspaceMember add a WorkspaceMember to a Workspace
func (repo *WorkspaceRepo) AddWorkspaceMember(workspaceID uuid.UUID, userID uuid.UUID) (*models.WorkspaceMember, error) {
	userMember := models.WorkspaceMember{
		WorkspaceID: workspaceID,
		UserID:      userID,
		Role:        "member",
	}
	if err := repo.db.Create(&userMember).Error; err != nil {
		return &models.WorkspaceMember{}, err
	}
	return &userMember, nil
}

// GetWorkspaceMember for a given workspaceID, userID pair
func (repo *WorkspaceRepo) GetWorkspaceMember(workspaceID string, userID string) (*models.WorkspaceMember, error) {
	m := &models.WorkspaceMember{}
	where := &models.WorkspaceMember{
		WorkspaceID: uuid.MustParse(workspaceID),
		UserID:      uuid.MustParse(userID),
	}
	if err := repo.db.Where(where).First(&m).Error; err != nil {
		return &models.WorkspaceMember{}, err
	}
	return m, nil
}
