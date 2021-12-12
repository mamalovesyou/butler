package models

import (
	"database/sql/driver"
	"github.com/butlerhq/butler/api/services/connectors/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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

type Connector struct {
	BaseModel
	WorkspaceID uuid.UUID
	AuthScheme  AuthScheme `sql:"AuthSchemeEnum"`
	Provider    string
	ExpiresIn   time.Time
	Secret      *ConnectorSecret `gorm:"foreignKey:ConnectorID"`
}

func (c *Connector) TableName() string {
	return "connectors"
}

// ToPb return the workspace.UserMembers of a OrganizationMember
func (c *Connector) ToPb() *connectors.WorkspaceConnector {
	pb := &connectors.WorkspaceConnector{
		Id:          c.ID.String(),
		WorkspaceId: c.WorkspaceID.String(),
		Name:        c.Provider,
		CreatedAt:   timestamppb.New(c.CreatedAt),
		UpdatedAt:   timestamppb.New(c.UpdatedAt),
	}

	if !c.ExpiresIn.IsZero() {
		pb.ExpiresIn = timestamppb.New(c.ExpiresIn)
	}

	return pb
}

type ConnectorSecret struct {
	BaseModel
	ConnectorID uuid.UUID
	Value       string
}

func (c *ConnectorSecret) TableName() string {
	return "connector_secrets"
}
