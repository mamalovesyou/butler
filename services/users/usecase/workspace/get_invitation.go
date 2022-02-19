package workspace

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (svc *WorkspaceUsecase) GetInvitation(ctx context.Context, invitationID, token string) (*models.Invitation, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.GetInvitation")
	defer span.Finish()

	invitation, err := svc.InvitationRepo.GetInvitation(invitationID, token)
	if err != nil {
		logger.Error(ctx, "Unable to retrieve invitation", zap.Error(err), zap.String("invitationID", invitationID))
		return &models.Invitation{}, err
	}

	return &invitation, nil
}
