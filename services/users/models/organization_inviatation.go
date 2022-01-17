package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/butlerhq/butler/api/services/users/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrganizationInvitation struct {
	BaseModel
	OrganizationID uuid.UUID
	Email          string
	Token          string
	ExpiresAt      time.Time
}

// BeforeCreate will set a random token when creating a new OrganizationInvitation.
func (o *OrganizationInvitation) BeforeCreate(tx *gorm.DB) (err error) {
	o.BaseModel.BeforeCreate(tx)
	o.Token = uuid.New().String()
	return nil
}

func (u *OrganizationInvitation) TableName() string {
	return "organization_invitations"
}

// ToPb return the users.OrganizationInvitation of an Organization model
func (o *OrganizationInvitation) ToPb() *users.Invitation {
	return &users.Invitation{
		Id:        o.ID.String(),
		Email:     o.Email,
		ExpiresAt: timestamppb.New(o.CreatedAt),
		CreatedAt: timestamppb.New(o.UpdatedAt),
	}
}
