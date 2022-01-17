package workspace

import (
	"context"
	"time"

	"github.com/butlerhq/butler/services/users/repositories"

	"github.com/butlerhq/butler/services/users/models"

	"github.com/opentracing/opentracing-go"
)

func (uc *WorkspaceUsecase) BulkInviteOrganizationMember(ctx context.Context, organizationID string, emails []string) ([]models.OrganizationInvitation, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.BulkInviteOrganizationMember")
	defer span.Finish()
	var err error

	invites := []models.OrganizationInvitation{}

	tx := func(repo *repositories.InvitationRepo) error {
		expiresAt := time.Now().Add(time.Hour * 48)
		invites, err = repo.CreateOrganizationInvitations(organizationID, emails, expiresAt)
		return err
	}
	if err = uc.InvitationRepo.WithTransaction(tx); err != nil {
		return []models.OrganizationInvitation{}, err
	}

	return invites, nil
}
