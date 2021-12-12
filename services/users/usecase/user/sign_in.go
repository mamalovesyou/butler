package user

import (
	"context"
	"errors"
	user_err "github.com/butlerhq/butler/services/users/errors"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

// SignIn user a user with an email/passsword combination
func (uc *UserUsecase) SignIn(ctx context.Context, email, password string) (*models.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user_ucase.SignIn")
	defer span.Finish()
	user, err := uc.UserRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_err.ErrUnknownEmail
		} else {
			return nil, err
		}
	}
	if err = uc.UserRepo.VerifyPassword(user, password); err != nil {
		return nil, user_err.ErrInvalidPassword
	}

	return uc.authenticate(ctx, user)
}
