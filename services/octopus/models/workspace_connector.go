package models

import (
	"database/sql/driver"
	"time"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthScheme string

const (
	OAUTH2  AuthScheme = "OAUTH2"
	API_KEY AuthScheme = "API_KEY"
)

func (as *AuthScheme) Scan(value interface{}) error {
	*as = AuthScheme(value.(string))
	return nil
}

func (as AuthScheme) Value() (driver.Value, error) {
	return string(as), nil
}

func (as AuthScheme) String() string {
	return string(as)
}

func (as AuthScheme) ToPb() octopus.AuthType {
	value := octopus.AuthType_value[as.String()]
	return octopus.AuthType(value)
}

type WorkspaceConnector struct {
	BaseModel
	WorkspaceID   uuid.UUID  `gorm:"uniqueIndex:connectors_workspace_provider_idx;index:connectors_workspace_idx"`
	AuthScheme    AuthScheme `sql:"auth_scheme_enum"`
	Provider      string     `gorm:"uniqueIndex:connectors_workspace_provider_idx"`
	ExpiresIn     time.Time
	Secret        *ConnectorSecret `gorm:"foreignKey:ConnectorID"`
	AccountConfig *ConnectorConfig `gorm:"foreignKey:ConnectorID"`
}

func (c *WorkspaceConnector) TableName() string {
	return "workspace_connectors"
}

// ToPb return the workspace.UserMembers of a OrganizationMember
func (c *WorkspaceConnector) ToPb() *octopus.WorkspaceConnector {
	pb := &octopus.WorkspaceConnector{
		Id:          c.ID.String(),
		WorkspaceId: c.WorkspaceID.String(),
		Name:        c.Provider,
		AuthScheme:  c.AuthScheme.ToPb(),
		CreatedAt:   timestamppb.New(c.CreatedAt),
		UpdatedAt:   timestamppb.New(c.UpdatedAt),
	}

	if !c.ExpiresIn.IsZero() {
		pb.ExpiresIn = timestamppb.New(c.ExpiresIn)
	}

	if c.AccountConfig != nil {
		pb.AccountConfig = c.AccountConfig.ToPb()
	}

	return pb
}

// ToConnectorSecretPairPb
func (c *WorkspaceConnector) ToConnectorSecretPairPb() *octopus.ConnectorSecretPair {
	pb := &octopus.ConnectorSecretPair{
		Connector: c.ToPb(),
	}
	if c.Secret != nil {
		pb.Credentials = c.Secret.ToPb()
	}

	return pb
}
