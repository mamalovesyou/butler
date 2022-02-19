package catalog

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/butlerhq/butler/internal/airbyte/sources"

	google_ads "github.com/butlerhq/butler/internal/airbyte/sources/google-ads"

	linkedin_ads "github.com/butlerhq/butler/internal/airbyte/sources/linkedin-ads"

	"github.com/cenkalti/backoff/v4"

	"github.com/butlerhq/butler/services/octopus/config"

	"github.com/butlerhq/airbyte-client-go/airbyte"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataSource interface {
	GetName() string
	GetAuthScheme() sources.AuthScheme
	GetAirbyteConfig(config, secrets map[string]interface{}) (interface{}, error)
	BindAirbyteSource(sourceID string, icon string)
	GetStreamCatalog() string
	ToPb() *octopus.DataSource
}

type DataSourceList []DataSource

func (ls *DataSourceList) ToPb() *octopus.DataSourceList {
	tmp := make([]*octopus.DataSource, len(*ls))
	for i, source := range *ls {
		tmp[i] = source.ToPb()
	}
	return &octopus.DataSourceList{Sources: tmp}
}

type Catalog struct {
	mu                  sync.Mutex
	isSyncedWithAirbyte bool
	sourcesByName       map[string]DataSource
	sourcesByAirbyteID  map[string]DataSource
	destinationsByName  map[string]DataSource
	airbyteURL          string
}

func NewCatalog(cfg *config.SourcesConfig, airbyteURL string) *Catalog {
	// Init data-sources
	googleSource := google_ads.NewGoogleAdsSource(cfg.Google, cfg.RedirectURL)
	linkedinSource := linkedin_ads.NewLinkedinAdsSource(cfg.Linkedin, cfg.RedirectURL)

	return &Catalog{
		airbyteURL: airbyteURL,
		sourcesByName: map[string]DataSource{
			googleSource.GetName():   googleSource,
			linkedinSource.GetName(): linkedinSource,
		},
	}
}

func (catalog *Catalog) Init() error {
	bkoff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3)
	return backoff.Retry(catalog.SyncWithAirbyte, bkoff)
}

// SyncWithAirbyte fetches all airbyte data-sources and map airbyte source to those data-sources
func (catalog *Catalog) SyncWithAirbyte() error {
	ctx := context.Background()
	logger.Info(ctx, "Syncing airbyte data-sources...", zap.String("airbyteURL", catalog.airbyteURL))
	client, err := airbyte.NewClientWithResponses(catalog.airbyteURL)
	if err != nil {
		logger.Error(ctx, "Unable to create airbyte client", zap.Error(err))
		return err
	}

	resp, err := client.ListSourceDefinitionsWithResponse(ctx)
	if err != nil {
		logger.Error(ctx, "Unable to sync airbyte data-sources definitions", zap.Error(err))
		return err
	}
	catalog.mu.Lock()
	sourcesByAirbyteID := make(map[string]DataSource, len(catalog.sourcesByName))
	for _, source := range resp.JSON200.SourceDefinitions {

		if availableSource, ok := catalog.sourcesByName[source.Name]; ok {
			logger.Debug(ctx, "Sync data source", zap.String("name", source.Name))
			availableSource.BindAirbyteSource(source.SourceDefinitionId, *source.Icon)
			sourcesByAirbyteID[source.SourceDefinitionId] = availableSource
		}
	}
	catalog.sourcesByAirbyteID = sourcesByAirbyteID
	catalog.isSyncedWithAirbyte = true
	catalog.mu.Unlock()
	return nil
}

// ListAvailableSources returns a list of available connector
func (catalog *Catalog) ListAvailableSources() DataSourceList {
	tmp := make(DataSourceList, 0, len(catalog.sourcesByName))
	catalog.mu.Lock()
	for _, value := range catalog.sourcesByName {
		tmp = append(tmp, value)
	}
	catalog.mu.Unlock()
	return tmp
}

// GetSource returns a DataSource for a given airbyteSourceID
func (catalog *Catalog) GetSource(airbyteSourceID string) (DataSource, bool) {
	catalog.mu.Lock()
	connector, ok := catalog.sourcesByAirbyteID[airbyteSourceID]
	catalog.mu.Unlock()
	return connector, ok
}

// GetSourceByName returns a DataSource for a given name
func (catalog *Catalog) GetSourceByName(name string) (DataSource, bool) {
	catalog.mu.Lock()
	connector, ok := catalog.sourcesByName[name]
	catalog.mu.Unlock()
	return connector, ok
}

func (catalog *Catalog) ExchangeOAuthCode(ctx context.Context, airbyteSOurceID, code string) (*oauth2.Token, error) {
	logger.Debug(ctx, "About to exchange oauth code", zap.String("airbyteSOurceID", airbyteSOurceID))
	catalog.mu.Lock()
	source, ok := catalog.sourcesByAirbyteID[airbyteSOurceID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unknown connector bind to airbyte sourceID: %s", airbyteSOurceID))
	}
	catalog.mu.Unlock()

	// Verify this is an oauth2 source
	if source.GetAuthScheme() != sources.OAUTH2 {
		return nil, errors.New(fmt.Sprintf("This source dosen't support oauth 2 auth: %s", source.GetName()))
	}

	token, err := source.(sources.OAuthSource).ExchangeCode(ctx, code)
	if err != nil {
		logger.Error(ctx, "Unable to exchange oauth code", zap.Error(err), zap.String("provider", source.GetName()))
		return nil, status.Errorf(codes.InvalidArgument, "Unable to exchange oauth code")
	}
	return token, nil
}
