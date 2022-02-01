package repositories

import (
	"context"

	"go.uber.org/zap"

	"gorm.io/gorm/clause"

	"github.com/google/uuid"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/octopus/models"
	"gorm.io/gorm"
)

type ConnectorRepo struct {
	db *gorm.DB
}

func NewConnectorRepo(db *gorm.DB) *ConnectorRepo {
	return &ConnectorRepo{db: db.Debug()}
}

// UpsertOne create or update a WorkspaceConnector if it doesn't exist
func (repo *ConnectorRepo) UpsertOne(connector *models.WorkspaceConnector) (*models.WorkspaceConnector, error) {
	if err := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "workspace_id"}, {Name: "provider"}},
		UpdateAll: true,
	}).Create(connector).Error; err != nil {
		logger.Error(context.Background(), "Unable to upsert connector", zap.Error(err))
		return nil, err
	}
	return repo.FindByWorkspaceAndProvider(connector.WorkspaceID, connector.Provider)
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

// FindById retrieve a connector by id
func (repo *ConnectorRepo) FindById(id uuid.UUID) (*models.WorkspaceConnector, error) {
	connector := models.WorkspaceConnector{BaseModel: models.BaseModel{ID: id}}
	if err := repo.db.Debug().Preload(clause.Associations).Where(&connector).First(&connector).Error; err != nil {
		return nil, err
	}
	return &connector, nil
}

// FindByWorkspaceAndProvider belonging to a given WorkspaceID and provider name
func (repo *ConnectorRepo) FindByWorkspaceAndProvider(workspaceID uuid.UUID, provider string) (*models.WorkspaceConnector, error) {
	result := models.WorkspaceConnector{}
	if err := repo.db.Debug().Preload(clause.Associations).Where(&models.WorkspaceConnector{
		WorkspaceID: workspaceID,
		Provider:    provider,
	}).First(&result).Error; err != nil {
		logger.Error(context.Background(), "Unable to retrieve workspace_connector", zap.Error(err))
		return nil, err
	}
	return &result, nil
}

// FindById retrieve a connector by id
func (repo *ConnectorRepo) UpsertConnectorSecret(secret models.ConnectorSecret) (*models.WorkspaceConnector, error) {
	if err := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "connector_id"}},
		UpdateAll: true,
	}).Create(&secret).Error; err != nil {
		return nil, err
	}

	return repo.FindById(secret.ConnectorID)
}

// FindById retrieve a connector by id
func (repo *ConnectorRepo) UpsertConnectorConfig(config models.ConnectorConfig) (*models.WorkspaceConnector, error) {
	if err := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "connector_id"}},
		UpdateAll: true,
	}).Create(&config).Error; err != nil {
		return nil, err
	}

	return repo.FindById(config.ConnectorID)
}
