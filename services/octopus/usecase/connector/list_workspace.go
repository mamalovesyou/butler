package connector

import (
	"context"

	"github.com/butlerhq/butler/internal/errors"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/octopus/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (uc *ConnectorUsecase) ListWorkspaceConnectors(ctx context.Context, workspaceID string) ([]models.WorkspaceConnector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.ListWorkspaceConnectors")
	defer span.Finish()

	result, err := uc.ConnectorRepo.ListByWorkspace(workspaceID)
	if err != nil {
		logger.Error(ctx, "Failed to get workspace connectors list", zap.Error(err))
		return []models.WorkspaceConnector{}, errors.ErrInternal
	}

	return result, nil
}
