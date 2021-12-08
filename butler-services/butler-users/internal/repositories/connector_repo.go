package repositories

import (
	"context"
	"github.com/butlerhq/butler/butler-core/logger"
	"github.com/butlerhq/butler/butler-services/butler-users/internal/models"
	"gorm.io/gorm"
)

type ConnectorRepo struct {
	db *gorm.DB
}

func NewConnectorRepo(db *gorm.DB) *ConnectorRepo {
	return &ConnectorRepo{db: db}
}

// CreateOne and save it in database
func (repo *ConnectorRepo) CreateOne(connector *models.Connector) (*models.Connector, error) {
	if err := repo.db.Create(connector).Error; err != nil {
		return &models.Connector{}, err
	}
	return connector, nil
}

// ListConnector belonging to a given Workspace
func (repo *ConnectorRepo) ListByWorkspace(workspaceID string) ([]models.Connector, error) {
	result := []models.Connector{}
	if err := repo.db.Debug().Where("workspace_id = ?", workspaceID).Find(&result).Error; err != nil {
		logger.Error(context.Background(), "Failed to list connectors")
		return nil, err
	}
	return result, nil
}