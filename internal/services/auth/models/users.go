package models

import (
	"github.com/butlerhq/butler/internal/services/gen/auth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	BaseModel
	Email        string
	FirstName    string
	LastName     string
	HashPassword string
}

func (u *User) TableName() string {
	return "users"
}

// ToPb return the proto-rest equivalent of a user
func (u *User) ToPb() *auth.User {
	return &auth.User{
		ID:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}
