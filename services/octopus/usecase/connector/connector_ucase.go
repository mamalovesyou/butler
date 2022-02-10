package connector

import (
	"github.com/butlerhq/butler/services/octopus/repositories"
	"github.com/butlerhq/butler/services/octopus/sources"
	"gorm.io/gorm"
)

type ConnectorUsecase struct {
	ConnectorRepo *repositories.ConnectorRepo
	Catalog       *sources.Catalog
}

func NewConnectorUsecase(db *gorm.DB, catalog *sources.Catalog) *ConnectorUsecase {
	return &ConnectorUsecase{
		ConnectorRepo: repositories.NewConnectorRepo(db),
		Catalog:       catalog,
	}
}
