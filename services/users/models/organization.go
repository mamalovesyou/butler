package models

import (
	"github.com/butlerhq/butler/api/services/users/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Organization struct {
	BaseModel
	Name        string
	OwnerID     uuid.UUID
	Workspaces  []Workspace          `gorm:"foreignKey:OrganizationID"`
	UserMembers []OrganizationMember `gorm:"foreignKey:OrganizationID"`
}

func (u *Organization) TableName() string {
	return "organizations"
}

// ToPb return the workspace.Organization of an Organization model
func (o *Organization) ToPb() *users.Organization {
	pb := &users.Organization{
		Id:        o.ID.String(),
		OwnerId:   o.OwnerID.String(),
		Name:      o.Name,
		CreatedAt: timestamppb.New(o.CreatedAt),
		UpdatedAt: timestamppb.New(o.UpdatedAt),
	}

	members := make([]*users.UserMember, len(o.UserMembers))
	for i, m := range o.UserMembers {
		members[i] = m.ToPb()
	}
	pb.Members = members

	workspaces := make([]*users.Workspace, len(o.Workspaces))
	for i, w := range o.Workspaces {
		workspaces[i] = w.ToPb()
	}
	pb.Workspaces = workspaces

	return pb
}

type OrganizationMember struct {
	BaseModel
	UserID         uuid.UUID
	OrganizationID uuid.UUID
	Role           string
}

func (u *OrganizationMember) TableName() string {
	return "organization_members"
}

// ToPb return the workspace.UserMember of a OrganizationMember
func (m *OrganizationMember) ToPb() *users.UserMember {
	return &users.UserMember{
		UserId: m.UserID.String(),
		Role:   m.Role,
	}
}
