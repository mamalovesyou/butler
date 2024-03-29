package services

import (
	"context"

	"github.com/butlerhq/butler/internal/airbyte/sources/catalog"

	"github.com/butlerhq/butler/api/services/octopus/v1"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DataSourcesService struct {
	Catalog *catalog.Catalog
	octopus.DataSourcesServiceServer
}

func NewDataSourcesService(catalog *catalog.Catalog) *DataSourcesService {
	return &DataSourcesService{
		Catalog: catalog,
	}
}

func (svc *DataSourcesService) ListAvailableSources(ctx context.Context, req *emptypb.Empty) (*octopus.DataSourceList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "data-sources-svc.ListAvailableSources")
	defer span.Finish()
	sources := svc.Catalog.ListAvailableSources()
	return sources.ToPb(), nil
}

func (svc *DataSourcesService) RegisterGRPCServer(server *grpc.Server) {
	octopus.RegisterDataSourcesServiceServer(server, svc)
}
