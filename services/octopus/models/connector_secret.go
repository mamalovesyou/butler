package models

import (
	"encoding/json"

	"google.golang.org/protobuf/types/known/structpb"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/google/uuid"
)

type ConnectorSecret struct {
	BaseModel
	ConnectorID uuid.UUID `gorm:"uniqueIndex:connector_secrets_connector_idx"`
	Connector   Connector `gorm:"foreignKey:ConnectorID"`
	Value       string
}

func (c *ConnectorSecret) TableName() string {
	return "connector_secrets"
}

func (c *ConnectorSecret) GetValueMap() map[string]interface{} {
	valueMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(c.Value), &valueMap)
	if err != nil {
		panic(err)
	}
	return valueMap
}

func (c *ConnectorSecret) ToPb() *octopus.ConnectorSecret {
	result, err := structpb.NewStruct(c.GetValueMap())
	if err != nil {
		panic(err)
	}
	return &octopus.ConnectorSecret{Value: result}
}
