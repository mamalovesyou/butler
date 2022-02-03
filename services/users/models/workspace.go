package models

import (
	"github.com/butlerhq/butler/api/services/users/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Workspace struct {
	BaseModel
	Name               string
	Description        string
	OrganizationID     uuid.UUID
	UserMembers        []WorkspaceMember `gorm:"foreignKey:WorkspaceID"`
	PendingInvitations []Invitation      `gorm:"foreignKey:WorkspaceID"`
	AirbyteWorkspaceID string
}

func (u *Workspace) TableName() string {
	return "workspaces"
}

// ToPb return the workspace.UserMembers of a OrganizationMember
func (w *Workspace) ToPb() *users.Workspace {
	pb := &users.Workspace{
		Id:             w.ID.String(),
		OrganizationId: w.OrganizationID.String(),
		Description:    w.Description,
		Name:           w.Name,
		CreatedAt:      timestamppb.New(w.CreatedAt),
		UpdatedAt:      timestamppb.New(w.UpdatedAt),
	}

	members := make([]*users.UserMember, len(w.UserMembers))
	for i, m := range w.UserMembers {
		members[i] = m.ToPb()
	}
	pb.Members = members

	invitations := make([]*users.Invitation, len(w.PendingInvitations))
	for i, invite := range w.PendingInvitations {
		invitations[i] = invite.ToPb()
	}
	pb.Invitations = invitations

	return pb
}

type WorkspaceMember struct {
	BaseModel
	UserID      uuid.UUID
	User        User
	WorkspaceID uuid.UUID
	Role        string
}

func (u *WorkspaceMember) TableName() string {
	return "workspace_members"
}

// ToPb return the workspace.UserMembers of a WorkspaceMember
func (m *WorkspaceMember) ToPb() *users.UserMember {
	return &users.UserMember{
		UserId:    m.UserID.String(),
		FirstName: m.User.FirstName,
		LastName:  m.User.LastName,
		Role:      m.Role,
	}
}
