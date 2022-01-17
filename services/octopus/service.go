package octopus

import (
	api_ocotpus "github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/butlerhq/butler/services/octopus/usecase/catalog"
	"github.com/butlerhq/butler/services/octopus/usecase/connector"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// OctopusService has router and db instances
type OctopusService struct {
	CatalogUsecase   *catalog.CatalogUsecase
	ConnectorUsecase *connector.ConnectorUsecase

	api_ocotpus.UnimplementedOctopusServiceServer
}

// NewOctopusService initialize with predefined configuration
func NewOctopusService(cfg *ServiceConfig, db *gorm.DB) *OctopusService {
	return &OctopusService{
		CatalogUsecase:   catalog.NewCatalogUsecase(&cfg.Connectors),
		ConnectorUsecase: connector.NewConnectorUsecase(&cfg.Connectors, db),
	}
}

// RegisterGRPCServer Service to the specified grpc server
func (svc *OctopusService) RegisterGRPCServer(server *grpc.Server) {
	api_ocotpus.RegisterOctopusServiceServer(server, svc)
}
