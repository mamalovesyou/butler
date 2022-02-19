package google_ads

import (
	"encoding/json"

	"github.com/butlerhq/butler/internal/airbyte/sources"

	"github.com/butlerhq/butler/internal/airbyte"
)

type GoogleAdsAirbyteConfig struct {
	Credentials          *airbyte.Credentials `json:"credentials"`
	CustomerID           string               `json:"customer_id"`
	StartDate            string               `json:"start_date"`
	ConversionWindowDays int8                 `json:"conversion_window_days"`
}

func (gc *GoogleAdsSource) GetAirbyteConfig(config, secrets map[string]interface{}) (interface{}, error) {
	rawSecrets, err := json.Marshal(secrets)
	if err != nil {
		panic(err)
	}
	credentials := &airbyte.Credentials{
		AuthMethod:   "oAuth2.0",
		ClientID:     gc.OauthConfig.ClientID,
		ClientSecret: gc.OauthConfig.ClientSecret,
	}

	if err := json.Unmarshal(rawSecrets, credentials); err != nil {
		return nil, err
	}

	var airbyteConfig GoogleAdsAirbyteConfig
	rawConfig, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(rawConfig, &airbyteConfig); err != nil {
		return nil, err
	}
	airbyteConfig.Credentials = credentials

	// Verify config is valid
	if err = sources.ValidateConfig(GOOGLE_ADS_AIRBYTE_CONNECTION_SCHEMA, airbyteConfig); err != nil {
		return nil, err
	}

	return airbyteConfig, nil
}
