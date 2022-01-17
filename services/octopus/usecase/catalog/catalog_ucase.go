package catalog

import (
	"github.com/butlerhq/butler/services/octopus/connectors"
	"github.com/butlerhq/butler/services/octopus/repositories"
)

type CatalogUsecase struct {
	CatalogRepo *repositories.CatalogRepo
}

func NewCatalogUsecase(cfg *connectors.Config) *CatalogUsecase {
	return &CatalogUsecase{
		CatalogRepo: repositories.NewCatalogRepo(cfg),
	}
}
