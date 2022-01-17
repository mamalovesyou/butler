package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/butlerhq/butler/api/services/users/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkspaceInvitation struct {
	BaseModel
	WorkspaceID uuid.UUID
	Email       string
	Token       string
	ExpiresAt   time.Time
}

// BeforeCreate will set a random token when creating a new WorkspaceInvitation.
func (o *WorkspaceInvitation) BeforeCreate(tx *gorm.DB) (err error) {
	err = o.BaseModel.BeforeCreate(tx)
	o.Token = uuid.New().String()
	return nil
}

func (u *WorkspaceInvitation) TableName() string {
	return "workspace_invitations"
}

// ToPb return the users.WorkspaceInvitation of an Organization model
func (o *WorkspaceInvitation) ToPb() *users.Invitation {
	return &users.Invitation{
		Id:        o.ID.String(),
		Email:     o.Email,
		ExpiresAt: timestamppb.New(o.CreatedAt),
		CreatedAt: timestamppb.New(o.UpdatedAt),
	}
}
