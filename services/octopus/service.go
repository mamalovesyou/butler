package octopus

import (
	api_ocotpus "github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/butlerhq/butler/internal/airbyte/sources/catalog"
	"github.com/butlerhq/butler/services/octopus/usecase/connector"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// OctopusService has router and db instances
type OctopusService struct {
	ConnectorUsecase *connector.ConnectorUsecase
	api_ocotpus.UnimplementedOctopusServiceServer
}

// NewOctopusService initialize with predefined configuration
func NewOctopusService(db *gorm.DB, catalog *catalog.Catalog) *OctopusService {
	return &OctopusService{
		ConnectorUsecase: connector.NewConnectorUsecase(db, catalog),
	}
}

// RegisterGRPCServer Service to the specified grpc server
func (svc *OctopusService) RegisterGRPCServer(server *grpc.Server) {
	api_ocotpus.RegisterOctopusServiceServer(server, svc)
}
