package linkedin_ads

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/butlerhq/butler/internal/airbyte/sources"

	"github.com/pkg/errors"

	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"

	"github.com/butlerhq/butler/internal/airbyte"

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
				Name:                  LINKEDIN_ADS_AIRBYTE_NAME,
				ConfigInputJSONSchema: LINKEDIN_ADS_CONFIG_INPUT_JSON_SCHEMA,
				SyncCatalogJSON:       LINKEDIN_ADS_STREAMS_CATALOG,
				AuthScheme:            sources.OAUTH2,
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

type LinkedinAdsAirbyteConfig struct {
	AccountIDs  []int64              `json:"account_ids"`
	Credentials *airbyte.Credentials `json:"credentials"`
	StartDate   string               `json:"start_date"`
}

func (lc *LinkedinAdsSource) GetAirbyteConfig(config, secrets []byte) (interface{}, error) {
	var airbyteConfig LinkedinAdsAirbyteConfig
	if err := json.Unmarshal(config, &airbyteConfig); err != nil {
		return nil, err
	}

	credentials := &airbyte.Credentials{
		AuthMethod:   "oAuth2.0",
		ClientID:     lc.OauthConfig.ClientID,
		ClientSecret: lc.OauthConfig.ClientSecret,
	}
	if err := json.Unmarshal(secrets, credentials); err != nil {
		return nil, err
	}
	airbyteConfig.Credentials = credentials
	logger.Debug(context.Background(), "Airbyte config", zap.Any("config", airbyteConfig))
	return airbyteConfig, nil
}

func (lc *LinkedinAdsSource) ValidateAndFormatConfig(config map[string]interface{}) (map[string]interface{}, error) {
	formattedConfig := make(map[string]interface{})

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

func (lc *LinkedinAdsSource) GetStreamCatalog() string {
	return lc.SyncCatalogJSON
}
