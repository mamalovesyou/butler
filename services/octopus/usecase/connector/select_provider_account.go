package connector

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/butlerhq/butler/services/octopus/models"

	"github.com/butlerhq/butler/api/services/octopus/v1"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (uc *ConnectorUsecase) SelectProviderAccount(ctx context.Context, req *octopus.SelectAccountRequest) (*octopus.WorkspaceConnector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.SelectProviderAccount")
	defer span.Finish()

	fmt.Printf("Request payload: %v\n", req)
	accountConfig := models.ConnectorConfig{
		ConnectorID: uuid.MustParse(req.WorkspaceConnectorId),
		AccountID:   req.AccountId,
		AccountName: req.AccountName,
		IsTest:      req.IsTestAccount,
	}

	wsConnector, err := uc.ConnectorRepo.UpsertConnectorConfig(accountConfig)
	if err != nil {
		logger.Error(ctx, "Unable to find a workspace connector", zap.Error(err), zap.Any("request", req))
		return &octopus.WorkspaceConnector{}, err
	}

	return wsConnector.ToPb(), err
}
