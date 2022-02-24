package models

import (
	"github.com/butlerhq/butler/internal/airbyte/sources"
	"github.com/butlerhq/butler/internal/postgres/types"
	"gorm.io/gorm"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Connector struct {
	BaseModel
	WorkspaceID uuid.UUID          `gorm:"uniqueIndex:connectors_workspace_provider_idx;index:connectors_workspace_idx"`
	AuthScheme  sources.AuthScheme `sql:"auth_scheme_enum"`
	Provider    string             `gorm:"uniqueIndex:connectors_workspace_provider_idx"`
	Secret      *ConnectorSecret   `gorm:"foreignKey:ConnectorID"`
	Config      *types.JSONB       `gorm:"type:jsonb"`

	// Airbyte binding
	AirbyteSourceDefinitionID string
	AirbyteSourceID           string
	AirbyteDestinationID      string
	AirbyteConnectionID       string
	AirbyteWorkspaceID        string
	IsActive                  bool
}

func (c *Connector) TableName() string {
	return "connectors"
}

// BeforeCreate will set the config if missing
func (c *Connector) BeforeCreate(tx *gorm.DB) (err error) {
	if c.Config == nil {
		c.Config = &types.JSONB{}
	}
	return nil
}

// ToPb return the workspace.UserMembers of a OrganizationMember
func (c *Connector) ToPb() *octopus.Connector {
	pb := &octopus.Connector{
		Id:                        c.ID.String(),
		WorkspaceId:               c.WorkspaceID.String(),
		AirbyteSourceDefinitionId: c.AirbyteSourceDefinitionID,
		Name:                      c.Provider,
		AuthScheme:                c.AuthScheme.ToPb(),
		UpdatedAt:                 timestamppb.New(c.UpdatedAt),
		Config:                    c.Config.ToPbStruct(),
		IsActive:                  c.IsActive,
	}

	return pb
}
