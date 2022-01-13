package connector

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/octopus/models"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (uc *ConnectorUsecase) ConnectWithCode(ctx context.Context, workspaceID, provider, code string) (*models.WorkspaceConnector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connector_ucase.ConnectWithCode")
	defer span.Finish()

	fmt.Println("connectWithCode", workspaceID, provider, code)
	fmt.Println(uc.CatalogRepo)

	token, err := uc.CatalogRepo.ExchangeOAuthCode(ctx, provider, code)
	if err != nil {
		return &models.WorkspaceConnector{}, err
	}

	secretData, err := json.Marshal(token)
	if err != nil {
		logger.Error(ctx, "Unable to exchange oauth code", zap.Error(err), zap.String("provider", provider))
		return nil, err
	}

	connector := &models.WorkspaceConnector{
		WorkspaceID: uuid.MustParse(workspaceID),
		Provider:    provider,
		AuthScheme:  models.OAUTH2,
		ExpiresIn:   token.Expiry,
		Secret: &models.ConnectorSecret{
			Value: string(secretData),
		},
	}

	result, err := uc.ConnectorRepo.CreateOne(connector)
	if err != nil {
		logger.Error(ctx, "Failed to create workspace connector", zap.Error(err))
		return &models.WorkspaceConnector{}, err
	}

	return result, nil
}
