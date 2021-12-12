package user

import (
	"context"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/services/users/errors"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
)

// SignUp creates a new user
func (uc *UserUsecase) SignUp(ctx context.Context, firstname, lastname, email, password string) (*models.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user_ucase.SignUp")
	defer span.Finish()
	user := &models.User{
		FirstName:    firstname,
		LastName:     lastname,
		Email:        email,
		HashPassword: uc.UserRepo.HashPassword(password),
	}
	var err error
	if user, err = uc.UserRepo.CreateOne(user); err != nil {
		if postgres.IsDuplicateKeyError(err) {
			return nil, errors.ErrEmailAlreadyUsed
		} else {
			return nil, err
		}
	}

	return uc.authenticate(ctx, user)
}
