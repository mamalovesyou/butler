package linkedin_ads

import (
	"context"

	"github.com/butlerhq/butler/internal/airbyte/sources"

	"github.com/butlerhq/butler/services/octopus/config"

	"github.com/butlerhq/butler/api/services/octopus/v1"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

const LINKEDIN_ADS_AIRBYTE_NAME = "LinkedIn Ads"

type LinkedinAdsSource struct {
	sources.OAuth2DataSource
}

func NewLinkedinAdsSource(cfg config.OAuthSourceConfig, redirectURL string) *LinkedinAdsSource {
	oauthCfg := oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       []string{"r_liteprofile", "r_emailaddress", "rw_ads"},
		Endpoint:     linkedin.Endpoint,
		RedirectURL:  redirectURL,
	}

	return &LinkedinAdsSource{
		sources.OAuth2DataSource{
			BaseDataSource: sources.BaseDataSource{
				Name:                   LINKEDIN_ADS_AIRBYTE_NAME,
				ConfigInputJSONSchema:  LINKEDIN_ADS_CONFIG_INPUT_JSON_SCHEMA,
				SecretsInputJSONSchema: LINKEDIN_ADS_SECRETS_INPUT_JSON_SCHEMA,
				SyncCatalogJSON:        LINKEDIN_ADS_STREAMS_CATALOG,
				AuthScheme:             sources.OAUTH2,
			},
			OauthConfig: oauthCfg,
			AuthURL:     oauthCfg.AuthCodeURL("", oauth2.AccessTypeOffline, oauth2.ApprovalForce),
		}}
}

func (lc *LinkedinAdsSource) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := lc.OauthConfig.Exchange(ctx, code, oauth2.AccessTypeOffline)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (lc *LinkedinAdsSource) GetConfig() *oauth2.Config {
	return &lc.OauthConfig
}

func (lc *LinkedinAdsSource) ToPb() *octopus.DataSource {
	return lc.OAuth2DataSource.ToPb()
}

func (lc *LinkedinAdsSource) GetStreamCatalog() string {
	return lc.SyncCatalogJSON
}
