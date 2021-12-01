package usecases

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/matthieuberger/butler/internal/logger"
	"github.com/matthieuberger/butler/internal/services/gen/connectors"
	"github.com/matthieuberger/butler/internal/services/workspace/models"
	"github.com/matthieuberger/butler/internal/services/workspace/repositories"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ConnectorsUseCase struct {
	ConnectorRepo *repositories.ConnectorRepo
	connectors.UnimplementedConnectorsServiceServer
}

func NewConnectorsUseCase(db *gorm.DB) *ConnectorsUseCase {
	return &ConnectorsUseCase {
		ConnectorRepo: repositories.NewConnectorRepo(db),
	}
}


func (uc *ConnectorsUseCase) ListCatalogConnectors(ctx context.Context, req *connectors.WorkspaceRequest) (*connectors.CatalogConnectorList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors.ListCatalogConnectors")
	defer span.Finish()

	google := &connectors.CatalogConnector{
		Id: uuid.NewString(),
		Name: "google",
		AuthType: "oauth2",
		AuthUrl: GoogleConnectorConfig.AuthCodeURL("aradomstate"),
	}
	linkedin := &connectors.CatalogConnector{
		Id: uuid.NewString(),
		Name: "linkedin",
		AuthType: "oauth2",
		AuthUrl: LinkedInConnectorConfig.AuthCodeURL("aradomstate"),
	}

	return &connectors.CatalogConnectorList{
		Connectors: []*connectors.CatalogConnector{
			google, linkedin,
		},
	}, nil
}

func (uc *ConnectorsUseCase) ListWorkspaceConnectors(ctx context.Context, req *connectors.WorkspaceRequest) (*connectors.WorkspaceConnectorList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors.ListWorkspaceConnectors")
	defer span.Finish()

	connectorsList, err := uc.ConnectorRepo.ListByWorkspace(req.WorkspaceId)
	if err != nil {
		logger.Error(ctx, "Failed to get connectors list", zap.Error(err))
		return &connectors.WorkspaceConnectorList{}, ErrInternal
	}

	result := make([]*connectors.WorkspaceConnector, len(connectorsList))
	for i, connector := range connectorsList {
		logger.Debug(ctx, "Converting connector to pb", zap.Any("connector", connector))
		result[i] = connector.ToPb()
	}

	return &connectors.WorkspaceConnectorList{
		Connectors: result,
	}, nil
}

func (uc *ConnectorsUseCase) GetOauthConnectorAuthorization(ctx context.Context, req *connectors.OAuthAuthorizationRequest) (*connectors.WorkspaceConnector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "connectors.GetOauthConnectorAuthorization")
	defer span.Finish()

	token, err := ExchangeOAuthCode(ctx, req.Code.Name, req.Code.Code)
	if err != nil {
		return &connectors.WorkspaceConnector{}, err
	}

	data, err := json.Marshal(token)
	if err != nil {
		logger.Error(ctx, "Failed to marshal oauth token", zap.Error(err))
	}

	connector := &models.Connector{
		WorkspaceID: uuid.MustParse(req.WorkspaceId),
		Provider: req.Code.Name,
		AuthScheme: models.OAUTH2,
		ExpiresIn: token.Expiry,
		Secret: &models.ConnectorSecret{
			Value: string(data),
		},
	}

	result, err := uc.ConnectorRepo.CreateOne(connector)
	if err != nil {
		logger.Error(ctx, "Failed to create connector", zap.Error(err))
		return &connectors.WorkspaceConnector{}, ErrInternal
	}

	return result.ToPb(), nil
}