package catalog

import (
	"context"
	"fmt"

	"github.com/butlerhq/butler/services/octopus/models"

	"github.com/opentracing/opentracing-go"
)

func (uc *CatalogUsecase) GetCatalog(ctx context.Context) ([]models.CatalogConnector, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "catalog_ucase.GetCatalog")
	defer span.Finish()

	fmt.Println("GetCatalog")
	catalog := uc.CatalogRepo.ListAvailableConnectors()
	return catalog, nil
}
