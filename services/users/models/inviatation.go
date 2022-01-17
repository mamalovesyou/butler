package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/butlerhq/butler/api/services/users/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Invitation struct {
	BaseModel
	OrganizationID uuid.UUID
	WorkspaceID    uuid.UUID
	Organization   Organization `gorm:"foreignKey:ID"`
	Workspace      Workspace    `gorm:"foreignKey:ID"`
	Email          string
	Token          string
	ExpiresAt      time.Time
}

// BeforeCreate will set a random token when creating a new OrganizationInvitation.
func (o *Invitation) BeforeCreate(tx *gorm.DB) (err error) {
	o.BaseModel.BeforeCreate(tx)
	o.Token = uuid.New().String()
	return nil
}

func (u *Invitation) TableName() string {
	return "invitations"
}

// ToPb return the users.OrganizationInvitation of an Organization model
func (o *Invitation) ToPb() *users.Invitation {
	return &users.Invitation{
		Id:           o.ID.String(),
		Email:        o.Email,
		Organization: o.Organization.ToPb(),
		Workspace:    o.Workspace.ToPb(),
		ExpiresAt:    timestamppb.New(o.CreatedAt),
	}
}
