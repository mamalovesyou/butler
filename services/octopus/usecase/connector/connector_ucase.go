package connector

import (
	"github.com/butlerhq/butler/services/octopus/repositories"
	"gorm.io/gorm"
)

type ConnectorUsecase struct {
	ConnectorRepo *repositories.ConnectorRepo
	CatalogRepo   *repositories.CatalogRepo
}

func NewConnectorUsecase(db *gorm.DB) *ConnectorUsecase {
	return &ConnectorUsecase{
		ConnectorRepo: repositories.NewConnectorRepo(db),
	}
}
