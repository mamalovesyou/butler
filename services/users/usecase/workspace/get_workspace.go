package workspace

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (svc *WorkspaceUsecase) GetWorkspace(ctx context.Context, workspaceID string) (*models.Workspace, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace_ucase.GetWorkspace")
	defer span.Finish()

	// TODO: Check permissions

	workspace, err := svc.WorkspaceRepo.FindByID(workspaceID)
	if err != nil {
		logger.Error(ctx, "Unable to retrieve workspace", zap.Error(err))
		return nil, err
	}

	return workspace, nil
}
