package google_ads

import (
	"context"

	"github.com/butlerhq/butler/internal/airbyte/sources"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/butlerhq/butler/services/octopus/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const GOOGLE_ADS_AIRBYTE_NAME = "Google Ads"

type GoogleAdsSource struct {
	sources.OAuth2DataSource
}

func NewGoogleAdsSource(cfg config.OAuthSourceConfig, redirectURL string) *GoogleAdsSource {
	oauthCfg := oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/user.organization.read",
			"https://www.googleapis.com/auth/adwords",
		},
		Endpoint:    google.Endpoint,
		RedirectURL: redirectURL,
	}
	return &GoogleAdsSource{
		sources.OAuth2DataSource{
			BaseDataSource: sources.BaseDataSource{
				Name:                   GOOGLE_ADS_AIRBYTE_NAME,
				ConfigInputJSONSchema:  GOOGLE_ADS_CONFIG_INPUT_JSON_SCHEMA,
				SecretsInputJSONSchema: GOOGLE_ADS_SECRETS_INPUT_JSON_SCHEMA,
				SyncCatalogJSON:        GOOGLE_ADS_STREAMS_CATALOG,
				AuthScheme:             sources.OAUTH2,
			},
			OauthConfig: oauthCfg,
			AuthURL:     oauthCfg.AuthCodeURL("", oauth2.AccessTypeOffline, oauth2.ApprovalForce),
		}}
}

func (gc *GoogleAdsSource) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := gc.OauthConfig.Exchange(ctx, code, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (gc *GoogleAdsSource) GetStreamCatalog() string {
	return gc.SyncCatalogJSON
}

func (gc *GoogleAdsSource) GetConfig() *oauth2.Config {
	return &gc.OauthConfig
}

func (gc *GoogleAdsSource) ToPb() *octopus.DataSource {
	return gc.OAuth2DataSource.ToPb()
}
