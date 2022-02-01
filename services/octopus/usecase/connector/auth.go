package connector

import (
	"context"
	"encoding/json"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/octopus/models"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (uc *ConnectorUsecase) ConnectWithCode(ctx context.Context, workspaceID, provider, code string) (*models.WorkspaceConnector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connector_ucase.ConnectWithCode")
	defer span.Finish()

	token, err := uc.CatalogRepo.ExchangeOAuthCode(ctx, provider, code)
	if err != nil {
		logger.Error(ctx, "Unable to exchange oauth code", zap.Error(err))
		return &models.WorkspaceConnector{}, err
	}

	secretData, err := json.Marshal(token)
	if err != nil {
		logger.Error(ctx, "Unable to exchange oauth code", zap.Error(err), zap.String("provider", provider))
		return nil, err
	}
	connector := models.WorkspaceConnector{
		WorkspaceID: uuid.MustParse(workspaceID),
		Provider:    provider,
		AuthScheme:  models.OAUTH2,
		ExpiresIn:   token.Expiry,
	}
	if connector, err := uc.ConnectorRepo.UpsertOne(&connector); err != nil {
		logger.Error(ctx, "Failed to create workspace connector", zap.Error(err))
		return &models.WorkspaceConnector{}, err
	} else {
		logger.Debug(ctx, "About to update secret of connector", zap.Any("connector", connector))
		if connector, err = uc.ConnectorRepo.UpsertConnectorSecret(models.ConnectorSecret{
			ConnectorID: connector.ID,
			Value:       string(secretData),
		}); err != nil {
			logger.Error(ctx, "Failed to set connector secret", zap.Error(err))
			return &models.WorkspaceConnector{}, err
		}
		return connector, nil
	}
}
