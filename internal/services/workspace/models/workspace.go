package models

import (
	"github.com/google/uuid"
	"github.com/butlerhq/butler/internal/services/gen/workspace"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Workspace struct {
	BaseModel
	Name           string
	Description    string
	OrganizationID uuid.UUID
	UserMembers    []WorkspaceMember `gorm:"foreignKey:WorkspaceID"`
}

func (u *Workspace) TableName() string {
	return "workspaces"
}

// ToPb return the workspace.UserMembers of a OrganizationMember
func (w *Workspace) ToPb() *workspace.Workspace {
	pb := &workspace.Workspace{
		ID:             w.ID.String(),
		OrganizationID: w.OrganizationID.String(),
		Description:    w.Description,
		Name:           w.Name,
		CreatedAt:      timestamppb.New(w.CreatedAt),
		UpdatedAt:      timestamppb.New(w.UpdatedAt),
	}
	members := make([]*workspace.UserMember, len(w.UserMembers))
	for i, m := range w.UserMembers {
		members[i] = m.ToPb()
	}
	pb.Members = members
	return pb
}

type WorkspaceMember struct {
	BaseModel
	UserID      uuid.UUID
	WorkspaceID uuid.UUID
	Role        string
}

func (u *WorkspaceMember) TableName() string {
	return "workspace_members"
}

// ToPb return the workspace.UserMembers of a WorkspaceMember
func (m *WorkspaceMember) ToPb() *workspace.UserMember {
	return &workspace.UserMember{
		UserID: m.UserID.String(),
		Role:   m.Role,
	}
}
