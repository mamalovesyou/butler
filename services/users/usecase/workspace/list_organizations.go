package workspace

import (
	"context"

	"github.com/butlerhq/butler/internal/errors"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (svc *WorkspaceUsecase) ListOrganizations(ctx context.Context, userID string) ([]models.Organization, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.ListOrganizations")
	defer span.Finish()

	orgList, err := svc.OrganizationRepo.ListByUserID(userID)
	if err != nil {
		logger.Error(ctx, "Failed to get organization list", zap.Error(err))
		return nil, errors.ErrInternal
	}

	return orgList, nil
}
