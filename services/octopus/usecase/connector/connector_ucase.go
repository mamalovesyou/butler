package connector

import (
	"github.com/butlerhq/butler/internal/airbyte/sources/catalog"
	"github.com/butlerhq/butler/services/octopus/repositories"
	"gorm.io/gorm"
)

type ConnectorUsecase struct {
	ConnectorRepo *repositories.ConnectorRepo
	Catalog       *catalog.Catalog
}

func NewConnectorUsecase(db *gorm.DB, catalog *catalog.Catalog) *ConnectorUsecase {
	return &ConnectorUsecase{
		ConnectorRepo: repositories.NewConnectorRepo(db),
		Catalog:       catalog,
	}
}
