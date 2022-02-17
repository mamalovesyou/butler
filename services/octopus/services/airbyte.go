package services

import (
	"context"

	"github.com/butlerhq/airbyte-client-go/airbyte"
	"github.com/butlerhq/butler/internal/errors"
	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"
)

type AirbyteService struct {
	AirbyteURL string
}

func NewAirbyteService(url string) *AirbyteService {
	return &AirbyteService{AirbyteURL: url}
}

func (svc *AirbyteService) CreateSource(ctx context.Context, req *airbyte.CreateSourceJSONRequestBody) (*airbyte.SourceRead, error) {
	logger.Info(ctx, "Syncing connector config with airbyte...")
	client, err := airbyte.NewClientWithResponses(svc.AirbyteURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return &airbyte.SourceRead{}, errors.ErrInternal
	}

	resp, err := client.CreateSourceWithResponse(ctx, *req)
	logger.Debug(ctx, "Airbyte svc create resp", zap.Any("response", resp.JSON422))
	if err != nil {
		logger.Error(ctx, "Unable to create source on airbyte", zap.Error(err))
		return &airbyte.SourceRead{}, errors.ErrInternal
	}
	if resp.JSON422 != nil {
		logger.Error(ctx, "Cannot create airbyte source, invalid arguments", zap.Any("error", resp.JSON422))
		return &airbyte.SourceRead{}, errors.ErrInvalidArguments
	}

	return resp.JSON200, nil
}

func (svc *AirbyteService) UpdateSource(ctx context.Context, req *airbyte.UpdateSourceJSONRequestBody) (*airbyte.SourceRead, error) {
	logger.Info(ctx, "Syncing connector config with airbyte...")
	client, err := airbyte.NewClientWithResponses(svc.AirbyteURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return &airbyte.SourceRead{}, errors.ErrInternal
	}

	resp, err := client.UpdateSourceWithResponse(ctx, *req)
	if err != nil {
		logger.Error(ctx, "Unable to update source on airbyte", zap.Error(err))
		return &airbyte.SourceRead{}, errors.ErrInternal
	}

	return resp.JSON200, nil
}

func (svc *AirbyteService) CheckConnectionForAProposedUpdate(ctx context.Context, req *airbyte.CheckConnectionToSourceForUpdateJSONRequestBody) (*airbyte.CheckConnectionRead, error) {
	logger.Info(ctx, "Syncing connector config with airbyte...")
	client, err := airbyte.NewClientWithResponses(svc.AirbyteURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return &airbyte.CheckConnectionRead{}, errors.ErrInternal
	}

	resp, err := client.CheckConnectionToSourceForUpdateWithResponse(ctx, *req)
	if err != nil {
		logger.Error(ctx, "Unable to check connection source on airbyte", zap.Error(err))
		return &airbyte.CheckConnectionRead{}, errors.ErrInternal
	}

	if resp.JSON422 != nil {
		return nil, errors.ErrInvalidArguments
	}
	if resp.JSON404 != nil {
		return nil, errors.ErrInternal
	}

	return resp.JSON200, nil
}

func (svc *AirbyteService) CheckConnection(ctx context.Context, req *airbyte.ExecuteSourceCheckConnectionJSONRequestBody) (*airbyte.CheckConnectionRead, error) {
	logger.Info(ctx, "Airbyte check connection...")
	client, err := airbyte.NewClientWithResponses(svc.AirbyteURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return &airbyte.CheckConnectionRead{}, errors.ErrInternal
	}
	logger.Debug(ctx, "With config", zap.Any("config", req.ConnectionConfiguration))
	resp, err := client.ExecuteSourceCheckConnectionWithResponse(ctx, *req)
	if err != nil {
		logger.Error(ctx, "Unable to check connection source on airbyte", zap.Error(err))
		return &airbyte.CheckConnectionRead{}, errors.ErrInternal
	}

	if resp.JSON422 != nil {
		return nil, errors.ErrInvalidArguments
	}

	return resp.JSON200, nil
}

func (svc *AirbyteService) CreateConnection(ctx context.Context, req *airbyte.CreateConnectionJSONRequestBody) (*airbyte.ConnectionRead, error) {
	logger.Info(ctx, "Airbyte create connection...")
	client, err := airbyte.NewClientWithResponses(svc.AirbyteURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return nil, errors.ErrInternal
	}
	resp, err := client.CreateConnectionWithResponse(ctx, *req)
	if err != nil {
		logger.Error(ctx, "Unable to create connection source on airbyte", zap.Error(err))
		return nil, errors.ErrInternal
	}

	if resp.JSON422 != nil {
		return nil, errors.NewInvalidArgsError(err)
	}

	return resp.JSON200, nil
}

func (svc *AirbyteService) UpdateConnection(ctx context.Context, req *airbyte.UpdateConnectionJSONRequestBody) (*airbyte.ConnectionRead, error) {
	logger.Info(ctx, "Airbyte update connection...")
	client, err := airbyte.NewClientWithResponses(svc.AirbyteURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return nil, errors.ErrInternal
	}
	resp, err := client.UpdateConnectionWithResponse(ctx, *req)
	if err != nil {
		logger.Error(ctx, "Unable to update connection source on airbyte", zap.Error(err))
		return nil, errors.ErrInternal
	}

	if resp.JSON422 != nil {
		return nil, errors.ErrInvalidArguments
	}

	return resp.JSON200, nil
}
