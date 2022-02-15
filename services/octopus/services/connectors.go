package services

import (
	"context"
	"encoding/json"

	airbyte2 "github.com/butlerhq/butler/internal/airbyte/sources/catalog"

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
	catalog    *airbyte2.Catalog
	db         *gorm.DB
	repo       *repositories.ConnectorRepo
	airbyteSvc *AirbyteService
	octopus.UnimplementedConnectorsServiceServer
}

func NewConnectorsService(catalog *airbyte2.Catalog, db *gorm.DB, airbyteSvc *AirbyteService) *ConnectorsService {
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

	// validate and format config if present
	mapConfig := req.Config.AsMap()
	if len(mapConfig) > 0 {
		validConfig, err := source.ValidateAndFormatConfig(req.Config.AsMap())
		if err != nil {
			return &octopus.Connector{}, butler_errors.ErrInvalidArguments
		}
		mapConfig = validConfig
	}
	config := types.JSONB(mapConfig)
	connector := &models.Connector{
		WorkspaceID:               uuid.MustParse(req.WorkspaceId),
		AirbyteWorkspaceID:        req.AirbyteWorkspaceId,
		AirbyteSourceDefinitionID: req.AirbyteSourceDefinitionId,
		Provider:                  source.GetName(),
		AuthScheme:                source.GetAuthScheme(),
		Config:                    &config,
	}
	if err := svc.repo.CreateOne(connector); err != nil {
		logger.Error(ctx, "Unable to create connector", zap.Error(err))
		return &octopus.Connector{}, err
	}

	return connector.ToPb(), nil
}

func (svc *ConnectorsService) MutateConnector(ctx context.Context, req *octopus.MutateConnectorRequest) (*octopus.Connector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors-svc.MutateConnector")
	defer span.Finish()

	connector, err := svc.repo.FindById(uuid.MustParse(req.ConnectorId))
	if err != nil {
		logger.Error(ctx, "Unable to fetch connector", zap.Error(err), zap.String("connectorID", req.ConnectorId))
		return &octopus.Connector{}, err
	}

	source, ok := svc.catalog.GetSource(connector.AirbyteSourceDefinitionID)
	if !ok {
		logger.Error(ctx, "Unable to map a source", zap.String("airbyteSourceDefinitionId", connector.AirbyteSourceDefinitionID))
		return &octopus.Connector{}, butler_errors.ErrInvalidArguments
	}

	// validate and format config
	validConfig, err := source.ValidateAndFormatConfig(req.Config.AsMap())
	if err != nil {
		return &octopus.Connector{}, butler_errors.ErrInvalidArguments
	}
	jsonConfig := types.JSONB(validConfig)
	connector, err = svc.repo.UpdateOne(uuid.MustParse(req.ConnectorId), models.Connector{Config: &jsonConfig})
	if err != nil {
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

func (svc *ConnectorsService) TestConnection(ctx context.Context, req *octopus.TestConnectionRequest) (*octopus.TestConnectionResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors.TestConnection")
	defer span.Finish()

	connector, err := svc.repo.FindById(uuid.MustParse(req.ConnectorId))
	if err != nil {
		logger.Error(ctx, "Unable to fetch connector", zap.Error(err), zap.String("connectorID", req.ConnectorId))
		return &octopus.TestConnectionResponse{}, err
	}

	source, ok := svc.catalog.GetSource(connector.AirbyteSourceDefinitionID)
	if !ok {
		logger.Error(ctx, "Unable to map a source", zap.String("airbyteSourceDefinitionId", connector.AirbyteSourceDefinitionID))
		return &octopus.TestConnectionResponse{}, butler_errors.ErrInternal
	}

	validConfig, err := source.ValidateAndFormatConfig(req.Config.AsMap())
	if err != nil {
		logger.Error(ctx, "Unable to validate config ", zap.Any("config", req.Config.AsMap()), zap.Error(err))
		return &octopus.TestConnectionResponse{}, butler_errors.ErrInvalidArguments
	}
	configBytes, err := json.Marshal(validConfig)
	if err != nil {
		logger.Error(ctx, "Unable convert config to byte array ", zap.Error(err))
		return &octopus.TestConnectionResponse{}, butler_errors.ErrInternal
	}

	connectionConfig, err := source.GetAirbyteConfig(configBytes, []byte(connector.Secret.Value))
	if err != nil {
		logger.Error(ctx, "Unable to get airbyte config from source", zap.Any("source", source.GetName()), zap.Error(err))
		return &octopus.TestConnectionResponse{}, butler_errors.ErrInternal
	}
	check, err := svc.airbyteSvc.CheckConnection(ctx, &airbyte.ExecuteSourceCheckConnectionJSONRequestBody{
		SourceDefinitionId:      connector.AirbyteSourceDefinitionID,
		ConnectionConfiguration: connectionConfig,
	})
	if err != nil {
		logger.Error(ctx, "Unable to check connection with airbyte", zap.Error(err), zap.Any("config", connectionConfig))
		return &octopus.TestConnectionResponse{}, butler_errors.ErrInternal
	}

	result := &octopus.TestConnectionResponse{Status: string(check.Status)}
	if check.Message != nil {
		result.Message = *check.Message
	}

	if check.Status == "failed" {
		logger.Error(ctx, "Failed to check connection", zap.Any("check", check))
		return result, butler_errors.ErrInvalidArguments
	}

	if err := svc.db.Transaction(func(tx *gorm.DB) error {
		jsonbCfg := types.JSONB(validConfig)
		connectorUpdates := models.Connector{
			IsActive: true,
			Config:   &jsonbCfg,
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
				return err
			}
			logger.Debug(ctx, "Airbyte source created", zap.Any("airbyteSource", airbyteSource))
			connectorUpdates.AirbyteSourceID = airbyteSource.SourceId
		}

		if len(connector.AirbyteConnectionID) == 0 {
			logger.Debug(ctx, "Missing AirbyteConnectionID. Creating connection...")
			// TODO: pass schedule from connector
			connConfig := connections.NewConnectionConfig(connector.AirbyteSourceID, connector.AirbyteDestinationID, source.GetStreamCatalog())
			airbyteConn, err := svc.airbyteSvc.CreateConnection(ctx, connConfig.ToAirbyteCreateConnectionRequestBody())
			if err != nil {
				logger.Error(ctx, "Unable to create airbyte connection", zap.Any("config", connConfig), zap.Error(err))
				return butler_errors.ErrInternal
			}
			connectorUpdates.AirbyteConnectionID = airbyteConn.ConnectionId
		}

		if connector, err = repositories.NewConnectorRepo(tx).UpdateOne(connector.ID, connectorUpdates); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return &octopus.TestConnectionResponse{}, err
	}

	return result, nil
}

// TODO: cleanup
//func (svc *ConnectorsService) SyncConnection(ctx context.Context, req *octopus.SyncConnectionRequest) (*emptypb.Empty, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors.SyncConnection")
//	defer span.Finish()
//
//	connector, err := svc.repo.FindById(uuid.MustParse(req.ConnectorId))
//	if err != nil {
//		logger.Error(ctx, "Unable to fetch connector", zap.Error(err), zap.String("connectorID", req.ConnectorId))
//		return &emptypb.Empty{}, err
//	}
//
//	source, ok := svc.catalog.GetSource(connector.AirbyteSourceDefinitionID)
//	if !ok {
//		logger.Error(ctx, "Unable to map a source", zap.String("airbyteSourceDefinitionId", connector.AirbyteSourceDefinitionID))
//		return &emptypb.Empty{}, butler_errors.ErrInternal
//	}
//
//	if len(connector.AirbyteConnectionID) == 0 {
//		logger.Debug(ctx, "Missing AirbyteConnectionID. Creating connection...")
//		connConfig := connections.NewConnectionConfig(connector.AirbyteSourceID, connector.AirbyteDestinationID, source.GetStreamCatalog())
//		airbyteConn, err := svc.airbyteSvc.CreateConnection(ctx, connConfig.ToAirbyteCreateConnectionRequestBody())
//		if err != nil {
//			logger.Error(ctx, "Unable to create airbyte connection", zap.Any("config", connConfig), zap.Error(err))
//			return &emptypb.Empty{}, butler_errors.ErrInternal
//		}
//		connector, err = svc.repo.UpdateOne(uuid.MustParse(req.ConnectorId), models.Connector{AirbyteConnectionID: airbyteConn.ConnectionId})
//		if err != nil {
//			logger.Error(ctx, "Unable to set connector AirbyteConnectionID", zap.Error(err))
//			return &emptypb.Empty{}, err
//		}
//	}
//
//	return &emptypb.Empty{}, nil
//}

func (svc *ConnectorsService) RegisterGRPCServer(server *grpc.Server) {
	octopus.RegisterConnectorsServiceServer(server, svc)
}
