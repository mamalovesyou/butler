package services

import (
	"context"
	"encoding/json"

	airbyte2 "github.com/butlerhq/butler/internal/airbyte"

	"github.com/butlerhq/butler/internal/utils"

	source_catalog "github.com/butlerhq/butler/internal/airbyte/sources/catalog"

	"github.com/butlerhq/butler/internal/airbyte/connections"

	"github.com/butlerhq/airbyte-client-go/airbyte"

	butler_errors "github.com/butlerhq/butler/internal/errors"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres/types"
	"go.uber.org/zap"

	"github.com/butlerhq/butler/services/octopus/models"
	"github.com/butlerhq/butler/services/octopus/repositories"
	"github.com/google/uuid"

	"gorm.io/gorm"

	"github.com/butlerhq/butler/api/services/octopus/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type ConnectorsService struct {
	catalog    *source_catalog.Catalog
	db         *gorm.DB
	repo       *repositories.ConnectorRepo
	airbyteSvc *AirbyteService
	octopus.UnimplementedConnectorsServiceServer
}

func NewConnectorsService(catalog *source_catalog.Catalog, db *gorm.DB, airbyteSvc *AirbyteService) *ConnectorsService {
	return &ConnectorsService{
		db:         db,
		catalog:    catalog,
		repo:       repositories.NewConnectorRepo(db),
		airbyteSvc: airbyteSvc,
	}
}

func (svc *ConnectorsService) CreateConnector(ctx context.Context, req *octopus.CreateConnectorRequest) (*octopus.Connector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors-svc.CreateConnector")
	defer span.Finish()

	source, ok := svc.catalog.GetSource(req.AirbyteSourceDefinitionId)
	if !ok {
		logger.Error(ctx, "Unable to map a source", zap.String("airbyteSourceDefinitionId", req.AirbyteSourceDefinitionId))
		return &octopus.Connector{}, butler_errors.ErrInvalidArguments
	}

	connector := &models.Connector{
		WorkspaceID:               uuid.MustParse(req.WorkspaceId),
		AirbyteWorkspaceID:        req.AirbyteWorkspaceId,
		AirbyteSourceDefinitionID: req.AirbyteSourceDefinitionId,
		Provider:                  source.GetName(),
		AuthScheme:                source.GetAuthScheme(),
	}
	if err := svc.repo.CreateOne(connector); err != nil {
		logger.Error(ctx, "Unable to create connector", zap.Error(err))
		return &octopus.Connector{}, err
	}

	return connector.ToPb(), nil
}

func (svc *ConnectorsService) GetConnector(ctx context.Context, req *octopus.GetConnectorRequest) (*octopus.Connector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors-svc.GetConnector")
	defer span.Finish()

	connector, err := svc.repo.FindById(uuid.MustParse(req.ConnectorId))
	if err != nil {
		logger.Error(ctx, "Unable to fetch connector", zap.Error(err), zap.String("connectorID", req.ConnectorId))
		return &octopus.Connector{}, err
	}

	return connector.ToPb(), nil
}

func (svc *ConnectorsService) AuthenticateOAuthConnector(ctx context.Context, req *octopus.AuthenticateConnectorRequest) (*octopus.Connector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors-svc.AuthenticateOAuthConnector")
	defer span.Finish()

	connector, err := svc.repo.FindById(uuid.MustParse(req.ConnectorId))
	if err != nil {
		logger.Error(ctx, "Unable to fetch connector", zap.Error(err), zap.String("connectorID", req.ConnectorId))
		return &octopus.Connector{}, err
	}

	token, err := svc.catalog.ExchangeOAuthCode(ctx, connector.AirbyteSourceDefinitionID, req.Code)
	if err != nil {
		logger.Error(ctx, "Unable to exchange oauth code", zap.Error(err), zap.String("connectorID", req.ConnectorId))
		return &octopus.Connector{}, err
	}

	secretData, err := json.Marshal(token)
	if err != nil {
		logger.Error(ctx, "Unable to unmarshal oauth token", zap.Error(err), zap.String("connectorID", req.ConnectorId))
		return &octopus.Connector{}, err
	}

	// Update connector secret
	if connector, err = svc.repo.UpsertConnectorSecret(models.ConnectorSecret{
		ConnectorID: connector.ID,
		Value:       string(secretData),
	}); err != nil {
		logger.Error(ctx, "Failed to update connector secret", zap.Error(err), zap.String("connectorID", req.ConnectorId))
		return &octopus.Connector{}, err
	}
	return connector.ToPb(), nil
}

func (svc *ConnectorsService) ListConnectors(ctx context.Context, req *octopus.ListConnectorsRequest) (*octopus.ConnectorList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.ConnectWithCode")
	defer span.Finish()

	connectors, err := svc.repo.ListByWorkspace(req.WorkspaceId)
	if err != nil {
		return &octopus.ConnectorList{}, err
	}

	result := make([]*octopus.Connector, len(connectors))
	for i, c := range connectors {
		logger.Debug(ctx, "Converting connector to pb", zap.Any("connector", c))
		result[i] = c.ToPb()
	}

	return &octopus.ConnectorList{
		Connectors: result,
	}, nil
}

func (svc *ConnectorsService) MutateConnector(ctx context.Context, req *octopus.MutateConnectorRequest) (*octopus.MutateConnectorResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors.TestConnection")
	defer span.Finish()

	connector, err := svc.repo.FindById(uuid.MustParse(req.ConnectorId))
	if err != nil {
		logger.Error(ctx, "Unable to fetch connector", zap.Error(err), zap.String("connectorID", req.ConnectorId))
		return &octopus.MutateConnectorResponse{}, err
	}

	source, ok := svc.catalog.GetSource(connector.AirbyteSourceDefinitionID)
	if !ok {
		logger.Error(ctx, "Unable to map a source", zap.String("airbyteSourceDefinitionId", connector.AirbyteSourceDefinitionID))
		return &octopus.MutateConnectorResponse{}, butler_errors.ErrInternal
	}

	// Merge secrets config
	secretsConfig := connector.Secret.MergeNewValues(req.Secrets.AsMap())

	// Verify config
	connectionConfig, err := source.GetAirbyteConfig(req.Config.AsMap(), secretsConfig)
	if err != nil {
		logger.Error(ctx, "Unable to get airbyte config from source", zap.Any("source", source.GetName()), zap.Error(err))
		return &octopus.MutateConnectorResponse{}, butler_errors.NewInvalidArgsError(err)
	}

	check, err := svc.airbyteSvc.CheckConnection(ctx, &airbyte.ExecuteSourceCheckConnectionJSONRequestBody{
		SourceDefinitionId:      connector.AirbyteSourceDefinitionID,
		ConnectionConfiguration: connectionConfig,
	})
	if err != nil {
		logger.Error(ctx, "Unable to check connection with airbyte", zap.Error(err), zap.Any("config", connectionConfig))
		return &octopus.MutateConnectorResponse{}, butler_errors.ErrInternal
	}

	result := &octopus.MutateConnectorResponse{
		Status:  string(check.Status),
		Message: utils.DerefString(check.Message),
		Logs:    check.JobInfo.Logs.LogLines,
	}

	// Skip updating connector if check connection failed
	if airbyte2.IsFailure(check) {
		logger.Debug(ctx, "Failed to check connection", zap.Any("check", check))
		return result, nil
	}

	jsonbCfg := types.JSONB(req.Config.AsMap())
	connectorUpdates := models.Connector{
		Config: &jsonbCfg,
	}

	if len(connector.AirbyteSourceID) == 0 {
		logger.Debug(ctx, "About to create resource on airbyte")
		airbyteSource, err := svc.airbyteSvc.CreateSource(ctx, &airbyte.CreateSourceJSONRequestBody{
			Name:                    connector.Provider,
			SourceDefinitionId:      connector.AirbyteSourceDefinitionID,
			WorkspaceId:             connector.AirbyteWorkspaceID,
			ConnectionConfiguration: connectionConfig,
		})
		if err != nil {
			logger.Error(ctx, "unable to create airbyte source", zap.Error(err))
			return &octopus.MutateConnectorResponse{}, butler_errors.ErrInternal
		}
		logger.Debug(ctx, "Airbyte source created", zap.Any("airbyteSource", airbyteSource))
		connectorUpdates.AirbyteSourceID = airbyteSource.SourceId
	} else {
		logger.Debug(ctx, "About to update resource on airbyte")
		_, err := svc.airbyteSvc.UpdateSource(ctx, &airbyte.UpdateSourceJSONRequestBody{
			SourceId:                connector.AirbyteSourceID,
			Name:                    connector.Provider,
			ConnectionConfiguration: connectionConfig,
		})
		if err != nil {
			logger.Error(ctx, "Unable to update airbyte source", zap.Error(err))
			return &octopus.MutateConnectorResponse{}, err
		}
	}

	// Perform connector update
	if connector, err = svc.repo.UpdateOne(connector.ID, connectorUpdates); err != nil {
		return &octopus.MutateConnectorResponse{}, err
	}

	// Create connection if missing
	if len(connector.AirbyteConnectionID) == 0 {
		logger.Debug(ctx, "Missing AirbyteConnectionID. Creating connection...")
		connConfig := connections.NewConnectionConfig(connector.AirbyteSourceID, connector.AirbyteDestinationID, source.GetStreamCatalog())
		body := connConfig.ToAirbyteCreateConnectionRequestBody()
		logger.Debug(ctx, "Request body", zap.Any("body", body))
		airbyteConn, err := svc.airbyteSvc.CreateConnection(ctx, body)
		if err != nil {
			logger.Error(ctx, "Unable to create airbyte connection", zap.Any("config", connConfig), zap.Error(err))
			return &octopus.MutateConnectorResponse{}, butler_errors.ErrInternal
		}
		connectorUpdates.AirbyteConnectionID = airbyteConn.ConnectionId
		if connector, err = svc.repo.UpdateOne(connector.ID, models.Connector{
			IsActive:            true,
			AirbyteConnectionID: airbyteConn.ConnectionId,
		}); err != nil {
			return &octopus.MutateConnectorResponse{}, err
		}
	}

	return result, nil
}

func (svc *ConnectorsService) RegisterGRPCServer(server *grpc.Server) {
	octopus.RegisterConnectorsServiceServer(server, svc)
}
