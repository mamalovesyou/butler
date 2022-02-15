package destinations

import (
	"context"
	"sync"

	"github.com/cenkalti/backoff/v4"

	"github.com/butlerhq/airbyte-client-go/airbyte"

	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"
)

type DestinationCatalog struct {
	mu                  sync.Mutex
	isSyncedWithAirbyte bool
	destinationsByName  map[string]Destination
	airbyteURL          string
}

func NewDestinationCatalog(airbyteURL string, destinations ...Destination) *DestinationCatalog {
	destinationsByName := make(map[string]Destination)
	for _, d := range destinations {
		destinationsByName[d.Name()] = d
	}

	return &DestinationCatalog{
		airbyteURL:         airbyteURL,
		destinationsByName: destinationsByName,
	}
}

func (catalog *DestinationCatalog) Init() error {
	bkoff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3)
	return backoff.Retry(catalog.SyncWithAirbyte, bkoff)
}

// SyncWithAirbyte fetches all airbyte destination and map them to allowed destination
func (catalog *DestinationCatalog) SyncWithAirbyte() error {
	ctx := context.Background()
	logger.Info(ctx, "Syncing airbyte data-sources...", zap.String("airbyteURL", catalog.airbyteURL))
	client, err := airbyte.NewClientWithResponses(catalog.airbyteURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return err
	}

	resp, err := client.ListDestinationDefinitionsWithResponse(ctx)
	if err != nil {
		logger.Error(ctx, "Unable to sync airbyte data-sources definitions", zap.Error(err))
		return err
	}
	catalog.mu.Lock()
	for _, dest := range resp.JSON200.DestinationDefinitions {
		if availableDestination, ok := catalog.destinationsByName[dest.Name]; ok {
			logger.Debug(ctx, "Sync airbyte destination", zap.String("name", dest.Name))
			availableDestination.BindAirbyteDefinition(dest.DestinationDefinitionId)
		}
	}
	catalog.isSyncedWithAirbyte = true
	catalog.mu.Unlock()
	return nil
}

// GetByName returns a DataSource for a given name
func (catalog *DestinationCatalog) GetByName(name string) (Destination, bool) {
	catalog.mu.Lock()
	connector, ok := catalog.destinationsByName[name]
	catalog.mu.Unlock()
	return connector, ok
}
