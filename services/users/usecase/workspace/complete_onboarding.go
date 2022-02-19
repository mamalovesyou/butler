package workspace

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (svc *WorkspaceUsecase) CompleteOnboarding(ctx context.Context, organizationID string) (*models.Organization, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.CompleteOnboarding")
	defer span.Finish()

	// TODO: verify user permissions Retrieve userID from context
	//userID, err := butlerctx.GetCtxTagUserID(ctx)
	//if err != nil {
	//	return nil, errors.ErrInvalidGrant
	//}

	updates := make(map[string]interface{})
	updates["onboarded"] = true
	org, err := svc.OrganizationRepo.UpdateOne(organizationID, updates)
	if err != nil {
		logger.Error(ctx, "Unable to update organization", zap.Error(err))
		return nil, err
	}

	return org, nil
}
