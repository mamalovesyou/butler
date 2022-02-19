package workspace

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (svc *WorkspaceUsecase) GetOrganization(ctx context.Context, organizationID string) (*models.Organization, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.GetOrganization")
	defer span.Finish()

	// TODO: Check permissions

	org, err := svc.OrganizationRepo.FindByID(organizationID)
	if err != nil {
		logger.Error(ctx, "Unable to retrieve organization", zap.Error(err))
		return nil, err
	}

	return org, nil
}
