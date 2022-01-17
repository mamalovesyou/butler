package workspace

import (
	"context"
	"time"

	"github.com/butlerhq/butler/services/users/repositories"

	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
)

func (uc *WorkspaceUsecase) BulkInviteWorkspaceMember(ctx context.Context, workspaceID string, emails []string) ([]models.WorkspaceInvitation, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.BulkInviteWorkspaceMember")
	defer span.Finish()
	var err error

	invites := []models.WorkspaceInvitation{}

	tx := func(repo *repositories.InvitationRepo) error {
		expiresAt := time.Now().Add(time.Hour * 48)
		invites, err = repo.CreateWorkspaceInvitations(workspaceID, emails, expiresAt)
		return err
	}
	if err = uc.InvitationRepo.WithTransaction(tx); err != nil {
		return []models.WorkspaceInvitation{}, err
	}

	return invites, nil
}
