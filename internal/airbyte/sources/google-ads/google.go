package google_ads

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/butlerhq/butler/internal/airbyte/sources"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"github.com/butlerhq/butler/services/octopus/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const GOOGLE_ADS_AIRBYTE_NAME = "Google Ads"

type GoogleAdWordsSource struct {
	sources.OAuth2DataSource
}

func NewGoogleAdWordsSource(cfg config.OAuthSourceConfig, redirectURL string) *GoogleAdWordsSource {
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
	return &GoogleAdWordsSource{
		sources.OAuth2DataSource{
			BaseDataSource: sources.BaseDataSource{
				Name:                  GOOGLE_ADS_AIRBYTE_NAME,
				ConfigInputJSONSchema: GOOGLE_ADS_CONFIG_INPUT_JSON_SCHEMA,
				SyncCatalogJSON:       GOOGLE_ADS_STREAMS_CATALOG,
				AuthScheme:            sources.OAUTH2,
			},
			OauthConfig: oauthCfg,
			AuthURL:     oauthCfg.AuthCodeURL("", oauth2.AccessTypeOffline, oauth2.ApprovalForce),
		}}
}

func (gc *GoogleAdWordsSource) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := gc.OauthConfig.Exchange(ctx, code, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (gc *GoogleAdWordsSource) ValidateAndFormatConfig(config map[string]interface{}) (map[string]interface{}, error) {
	var formattedConfig map[string]interface{}

	// account id parsing
	if accountID, ok := config["account_id"]; !ok {
		return nil, errors.New("InvalidConfig: Missing account_id in config")
	} else {
		accountIDInt64, err := strconv.ParseInt(fmt.Sprintf("%v", accountID), 10, 64)
		if err != nil {
			return nil, errors.New("InvalidConfig: account_id is invalid")
		}
		formattedConfig["account_ids"] = []int64{accountIDInt64}
	}

	// start date parsing
	if startDate, ok := config["start_date"]; !ok {
		return nil, errors.New("InvalidConfig: Missing start_date in config")
	} else {
		formattedConfig["start_date"] = fmt.Sprintf("%v", startDate)
	}

	return formattedConfig, nil
}

func (gc *GoogleAdWordsSource) GetStreamCatalog() string {
	return gc.SyncCatalogJSON
}

func (gc *GoogleAdWordsSource) GetConfig() *oauth2.Config {
	return &gc.OauthConfig
}

func (gc *GoogleAdWordsSource) ToPb() *octopus.DataSource {
	return gc.OAuth2DataSource.ToPb()
}
