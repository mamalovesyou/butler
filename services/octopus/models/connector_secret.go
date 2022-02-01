package models

import (
	"encoding/json"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/google/uuid"
)

type ConnectorSecret struct {
	BaseModel
	ConnectorID uuid.UUID          `gorm:"uniqueIndex:connector_secrets_connector_idx"`
	Connector   WorkspaceConnector `gorm:"foreignKey:ConnectorID"`
	Value       string
}

func (c *ConnectorSecret) TableName() string {
	return "connector_secrets"
}

func (c *ConnectorSecret) ToPb() *octopus.ConnectorSecret {
	valueMap := make(map[string]string)
	err := json.Unmarshal([]byte(c.Value), &valueMap)
	if err != nil {
		panic(err)
	}
	return &octopus.ConnectorSecret{Value: valueMap}
}
