package models

import (
	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/google/uuid"
)

type ConnectorConfig struct {
	BaseModel
	ConnectorID uuid.UUID
	Connector   WorkspaceConnector `gorm:"foreignKey:ConnectorID"`
	AccountID   string
	AccountName string
	IsTest      bool
}

func (m *ConnectorConfig) TableName() string {
	return "connector_configs"
}

// ToPb return the workspace.UserMembers of a OrganizationMember
func (m *ConnectorConfig) ToPb() *octopus.ProviderAccount {
	return &octopus.ProviderAccount{
		Id:   m.AccountID,
		Name: m.AccountName,
		Test: m.IsTest,
	}
}
