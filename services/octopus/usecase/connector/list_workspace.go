package connector

import (
	"context"

	"github.com/butlerhq/butler/internal/errors"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/octopus/models"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (uc *ConnectorUsecase) ListConnectors(ctx context.Context, workspaceID string) ([]models.Connector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.ListConnectors")
	defer span.Finish()

	result, err := uc.ConnectorRepo.ListByWorkspace(workspaceID)
	if err != nil {
		logger.Error(ctx, "Failed to get workspace data-sources list", zap.Error(err))
		return []models.Connector{}, errors.ErrInternal
	}

	return result, nil
}
