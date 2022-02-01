package connector

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/butlerhq/butler/services/octopus/models"

	"github.com/butlerhq/butler/internal/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/butlerhq/butler/api/services/octopus/v1"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

func (uc *ConnectorUsecase) ListAccounts(ctx context.Context, workspaceID, provider string) (*octopus.ListAccountsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.ListAccounts")
	defer span.Finish()
	connector, err := uc.ConnectorRepo.FindByWorkspaceAndProvider(
		uuid.MustParse(workspaceID),
		provider,
	)
	if err != nil {
		logger.Error(ctx, "Unable to find a workspace connector", zap.Error(err), zap.String("workspaceID", workspaceID), zap.String("provider", provider))
		return &octopus.ListAccountsResponse{}, err
	}
	logger.Debug(ctx, "Secrets value", zap.String("secret", connector.Secret.Value))

	var secrets models.ConnectorSecrets
	if err = json.Unmarshal([]byte(connector.Secret.Value), &secrets); err != nil {
		logger.Error(ctx, "Unable unmarshall secrets value", zap.Error(err), zap.String("workspaceID", workspaceID), zap.String("provider", provider))
		return &octopus.ListAccountsResponse{}, errors.ErrInternal
	}

	c, ok := uc.CatalogRepo.GetConnector(provider)
	if !ok {
		logger.Error(ctx, "Unable to find connector", zap.String("workspaceID", workspaceID), zap.String("provider", provider))
		return &octopus.ListAccountsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Unknown provider %s", provider))
	}

	return c.ListAccounts(ctx, &secrets)
}
