package octopus

import (
	"context"
	"fmt"

	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"

	"github.com/butlerhq/butler/api/services/octopus/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UseCase interface {
	RegisterUseCaseEndpoints(server *grpc.Server)
}

func (svc *OctopusService) GetCatalogConnectors(ctx context.Context, req *emptypb.Empty) (*octopus.CatalogConnectorList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.GetCatalogConnectors")
	defer span.Finish()

	catalog, err := svc.CatalogUsecase.GetCatalog(ctx)
	fmt.Println("Got catalog in delivrey: ")
	if err != nil {
		return &octopus.CatalogConnectorList{}, err
	}

	fmt.Println("Got catalog in delivrey: ")
	fmt.Printf("Catalog len: %d", len(catalog))

	result := make([]*octopus.CatalogConnector, len(catalog))
	for i, connector := range catalog {
		fmt.Println("Iterating catalog: ")
		fmt.Printf("Connector: %v", connector)
		result[i] = connector.ToPb()
	}

	return &octopus.CatalogConnectorList{
		Connectors: result,
	}, nil
}

func (svc *OctopusService) ConnectWithCode(ctx context.Context, req *octopus.ConnectWithCodeRequest) (*octopus.WorkspaceConnector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.ConnectWithCode")
	defer span.Finish()

	connector, err := svc.ConnectorUsecase.ConnectWithCode(ctx, req.WorkspaceId, req.Provider, req.Code)
	if err != nil {
		return &octopus.WorkspaceConnector{}, err
	}

	return connector.ToPb(), nil
}

func (svc *OctopusService) ListWorkspaceConnectors(ctx context.Context, req *octopus.WorkspaceConnectorsRequest) (*octopus.WorkspaceConnectorList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "octopus.ConnectWithCode")
	defer span.Finish()

	wsConnectors, err := svc.ConnectorUsecase.ListWorkspaceConnectors(ctx, req.WorkspaceId)
	if err != nil {
		return &octopus.WorkspaceConnectorList{}, err
	}

	result := make([]*octopus.WorkspaceConnector, len(wsConnectors))
	for i, connector := range wsConnectors {
		logger.Debug(ctx, "Converting connector to pb", zap.Any("connector", connector))
		result[i] = connector.ToPb()
	}

	return &octopus.WorkspaceConnectorList{
		Connectors: result,
	}, nil
}
