package models

import (
	"github.com/butlerhq/butler/api/services/users/v1"
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
func (u *User) ToPb() *users.User {
	return &users.User{
		ID:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

type AuthenticatedUser struct {
	User
	AccessToken  string
	RefreshToken string
}

// ToPb return the proto-rest equivalent of a user
func (u *AuthenticatedUser) ToPb() *users.AuthenticatedUser {
	return &users.AuthenticatedUser{
		User:         u.User.ToPb(),
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
	}
}
