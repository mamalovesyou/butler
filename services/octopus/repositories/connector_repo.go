package repositories

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/octopus/models"
	"gorm.io/gorm"
)

type ConnectorRepo struct {
	db *gorm.DB
}

func NewConnectorRepo(db *gorm.DB) *ConnectorRepo {
	return &ConnectorRepo{db: db}
}

// CreateOne and save it in database
func (repo *ConnectorRepo) CreateOne(connector *models.WorkspaceConnector) (*models.WorkspaceConnector, error) {
	if err := repo.db.Create(connector).Error; err != nil {
		return &models.WorkspaceConnector{}, err
	}
	return connector, nil
}

// ListConnector belonging to a given Workspace
func (repo *ConnectorRepo) ListByWorkspace(workspaceID string) ([]models.WorkspaceConnector, error) {
	result := []models.WorkspaceConnector{}
	if err := repo.db.Debug().Where("workspace_id = ?", workspaceID).Find(&result).Error; err != nil {
		logger.Error(context.Background(), "Failed to list octopus")
		return nil, err
	}
	return result, nil
}
