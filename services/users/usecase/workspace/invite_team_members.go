package workspace

import (
	"context"
	"time"

	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"

	"github.com/butlerhq/butler/services/users/sendgrid"
	"gorm.io/gorm"

	"github.com/butlerhq/butler/services/users/models"

	"github.com/opentracing/opentracing-go"
)

func (uc *WorkspaceUsecase) BatchInviteTeamMembers(ctx context.Context, organizationID string, workspaceID string, emails []string) ([]models.Invitation, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.BatchInviteTeamMembers")
	defer span.Finish()

	var err error

	invites := []models.Invitation{}

	if err = uc.InvitationRepo.DB().Transaction(func(tx *gorm.DB) error {
		var email *mail.SGMailV3
		_, err := uc.OrganizationRepo.FindByID(organizationID)
		if err != nil {
			logger.Error(ctx, "Unable to retrieve organization", zap.String("organizationID", organizationID))
			return err
		}

		expiresAt := time.Now().Add(time.Hour * 48)

		if len(workspaceID) > 0 {
			ws, err := uc.WorkspaceRepo.FindByID(workspaceID)
			if err != nil {
				logger.Error(ctx, "Unable to retrieve workspace", zap.String("workspaceID", workspaceID))
				return err
			}
			invites, err = uc.InvitationRepo.WithTransaction(tx).CreateWorkspaceInvitations(organizationID, workspaceID, emails, expiresAt)
			for _, invitation := range invites {
				email = sendgrid.NewInviteWorkspaceTeamMember(invitation.Email, ws.Name, uc.GetJoinLink(invitation.ID.String(), invitation.Token))
				uc.EmailClient.SendEmail(ctx, email)
			}
		} else {
			invites, err = uc.InvitationRepo.WithTransaction(tx).CreateWorkspaceInvitations(organizationID, workspaceID, emails, expiresAt)
			// TODO: Add missing template
			// mail := sendgrid.NewInviteWorkspaceTeamMember(emails, ws.Name, "link")
			//return uc.EmailClient.SendEmail(ctx, mail)
		}
		return nil
	}); err != nil {
		return []models.Invitation{}, err
	}

	return invites, nil
}
