package linkedin_ads

import (
	"encoding/json"

	"github.com/butlerhq/butler/internal/airbyte/sources"

	"github.com/butlerhq/butler/internal/airbyte"
)

type LinkedinAdsAirbyteConfig struct {
	AccountIDs  []int64              `json:"account_ids"`
	Credentials *airbyte.Credentials `json:"credentials"`
	StartDate   string               `json:"start_date"`
}

func (lc *LinkedinAdsSource) GetAirbyteConfig(config, secrets map[string]interface{}) (interface{}, error) {
	rawSecrets, err := json.Marshal(secrets)
	if err != nil {
		panic(err)
	}
	credentials := &airbyte.Credentials{
		AuthMethod:   "oAuth2.0",
		ClientID:     lc.OauthConfig.ClientID,
		ClientSecret: lc.OauthConfig.ClientSecret,
	}

	if err := json.Unmarshal(rawSecrets, credentials); err != nil {
		return nil, err
	}

	var airbyteConfig LinkedinAdsAirbyteConfig
	rawConfig, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(rawConfig, &airbyteConfig); err != nil {
		return nil, err
	}
	airbyteConfig.Credentials = credentials

	// Verify config is valid
	if err = sources.ValidateConfig(LINKEDIN_ADS_AIRBYTE_CONNECTION_SCHEMA, airbyteConfig); err != nil {
		return nil, err
	}

	return airbyteConfig, nil
}
