package repositories

import (
	"context"
	"errors"
	"fmt"

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

// CreateOne create a Connector if it doesn't exist
func (repo *ConnectorRepo) CreateOne(connector *models.Connector) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		var newConnector models.Connector
		result := tx.Where("workspace_id = ? AND provider = ?", connector.WorkspaceID, connector.Provider).First(&newConnector)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return tx.Create(connector).Error
			}
		} else {
			return errors.New(fmt.Sprintf("%s already exists in this workspace", connector.Provider))
		}
		return result.Error
	})
	return err
}

func (repo *ConnectorRepo) UpdateOne(connectorID uuid.UUID, updates models.Connector) (*models.Connector, error) {
	var connector models.Connector
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", connectorID.String()).First(&connector)
		if result.Error != nil {
			return result.Error
		}
		return tx.Model(&connector).Updates(updates).Error
	})
	return &connector, err
}

// UpsertOne create or update a Connector if it doesn't exist
func (repo *ConnectorRepo) UpsertOne(connector *models.Connector) (*models.Connector, error) {
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
func (repo *ConnectorRepo) ListByWorkspace(workspaceID string) ([]models.Connector, error) {
	result := []models.Connector{}
	if err := repo.db.Debug().Where("workspace_id = ?", workspaceID).Find(&result).Error; err != nil {
		logger.Error(context.Background(), "Failed to list connectors", zap.Error(err))
		return nil, err
	}
	return result, nil
}

// FindById retrieve a connector by id
func (repo *ConnectorRepo) FindById(id uuid.UUID) (*models.Connector, error) {
	connector := models.Connector{BaseModel: models.BaseModel{ID: id}}
	if err := repo.db.Debug().Preload(clause.Associations).Where(&connector).First(&connector).Error; err != nil {
		return nil, err
	}
	return &connector, nil
}

// FindByWorkspaceAndProvider belonging to a given WorkspaceID and provider name
func (repo *ConnectorRepo) FindByWorkspaceAndProvider(workspaceID uuid.UUID, provider string) (*models.Connector, error) {
	result := models.Connector{}
	if err := repo.db.Debug().Preload(clause.Associations).Where(&models.Connector{
		WorkspaceID: workspaceID,
		Provider:    provider,
	}).First(&result).Error; err != nil {
		logger.Error(context.Background(), "Unable to retrieve workspace_connector", zap.Error(err))
		return nil, err
	}
	return &result, nil
}

// FindById retrieve a connector by id
func (repo *ConnectorRepo) UpsertConnectorSecret(secret models.ConnectorSecret) (*models.Connector, error) {
	if err := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "connector_id"}},
		UpdateAll: true,
	}).Create(&secret).Error; err != nil {
		return nil, err
	}

	return repo.FindById(secret.ConnectorID)
}

// FindById retrieve a connector by id
func (repo *ConnectorRepo) UpsertConnectorConfig(config models.ConnectorConfig) (*models.Connector, error) {
	if err := repo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "connector_id"}},
		UpdateAll: true,
	}).Create(&config).Error; err != nil {
		return nil, err
	}

	return repo.FindById(config.ConnectorID)
}
