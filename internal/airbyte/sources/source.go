package sources

import (
	"context"
	"errors"
	"fmt"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"golang.org/x/oauth2"
)

type SyncCatalog struct {
	Streams []map[string]interface{} `json:"streams"`
}

type BaseDataSource struct {
	Name                      string
	AuthScheme                AuthScheme
	Icon                      string
	AirbyteSourceDefinitionID string
	ConfigInputJSONSchema     string
	SecretsInputJSONSchema    string
	SyncCatalogJSON           string
}

func (src *BaseDataSource) GetName() string {
	return src.Name
}

func (src *BaseDataSource) GetAuthScheme() AuthScheme {
	return src.AuthScheme
}

func (src *BaseDataSource) BindAirbyteSource(sourceDefinitionID, icon string) {
	src.AirbyteSourceDefinitionID = sourceDefinitionID
	src.Icon = icon
}

func (src *BaseDataSource) GetAirbyteConfig(config, secrets map[string]interface{}) (interface{}, error) {
	return nil, errors.New(fmt.Sprintf("%s doesn't implement GetAirbyteConfig", src.Name))
}

func (src *BaseDataSource) ToPb() *octopus.DataSource {
	authTypeInt := octopus.AuthType_value[string(src.AuthScheme)]
	return &octopus.DataSource{
		Name:                      src.Name,
		AuthType:                  octopus.AuthType(authTypeInt),
		IconSvg:                   src.Icon,
		ConfigInputJSONSchema:     src.ConfigInputJSONSchema,
		SecretsInputJSONSchema:    src.SecretsInputJSONSchema,
		AirbyteSourceDefinitionId: src.AirbyteSourceDefinitionID,
	}
}

type OAuthSource interface {
	ExchangeCode(context.Context, string) (*oauth2.Token, error)
}

type OAuth2DataSource struct {
	BaseDataSource
	OAuthSource
	AuthURL     string
	OauthConfig oauth2.Config
}

func (src *OAuth2DataSource) ExchangeCode(context.Context, string) (*oauth2.Token, error) {
	return nil, errors.New(fmt.Sprintf("Data source %s ExchangeCode is not implemented", src.Name))
}

func (src *OAuth2DataSource) ToPb() *octopus.DataSource {
	pbSource := src.BaseDataSource.ToPb()
	pbSource.AuthUrl = src.AuthURL
	return pbSource
}

type APIKeyDataSource struct {
	BaseDataSource
}
