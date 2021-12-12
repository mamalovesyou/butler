package workspace

import (
	"context"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/butlerhq/butler/services/users/services"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (svc *WorkspaceUsecase) CreateWorkspace(ctx context.Context, organizationID uuid.UUID, name, description string) (*models.Workspace, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.CreateWorkspace")
	defer span.Finish()

	// TODO: verify user is admin on ORGANIZATION_ID

	ws, err := svc.WorkspaceRepo.CreateOne(&models.Workspace{
		Name:           name,
		Description:    description,
		OrganizationID: organizationID,
	})

	if err != nil {
		logger.Error(ctx, "Unable to create workspace", zap.Error(err))
		return nil, services.ErrInternal
	}

	return ws, nil
}
