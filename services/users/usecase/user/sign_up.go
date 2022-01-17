package user

import (
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/services/users/errors"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
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

func (uc *UserUsecase) SignUpWithInvitation(ctx context.Context, firstname, lastname, password, invitationID, token string) (*models.AuthenticatedUser, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user_ucase.SignUpWithInvitation")
	defer span.Finish()

	invitation, err := uc.InvitationRepo.GetInvitation(invitationID, token)
	if err != nil {
		logger.Error(ctx, "Unable to retrieve invitation", zap.Error(err), zap.String("invitationID", invitationID))
		return &models.AuthenticatedUser{}, err
	}

	user := &models.User{
		FirstName:    firstname,
		LastName:     lastname,
		Email:        invitation.Email,
		HashPassword: uc.UserRepo.HashPassword(password),
	}

	err = uc.UserRepo.DB().Transaction(func(tx *gorm.DB) error {
		if user, err = uc.UserRepo.WithTransaction(tx).CreateOne(user); err != nil {
			return err
		}

		if invitation.WorkspaceID != uuid.Nil {
			_, err := uc.WorkspaceRepo.AddWorkspaceMember(invitation.WorkspaceID, user.ID)
			return err
		} else {
			_, err := uc.OrganizationRepo.AddOrganizationMember(invitation.OrganizationID, user.ID)
			return err
		}

		return nil
	})

	if postgres.IsDuplicateKeyError(err) {
		return &models.AuthenticatedUser{}, errors.ErrEmailAlreadyUsed
	}

	return uc.authenticate(ctx, user)
}
