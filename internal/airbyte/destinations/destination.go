package sources

import (
	"context"
	"errors"
	"fmt"

	"github.com/butlerhq/butler/api/services/octopus/v1"
	"golang.org/x/oauth2"
)

type BaseDataSource struct {
	name                      string
	authScheme                AuthScheme
	Icon                      string
	AirbyteSourceDefinitionID string
	ConfigInputJSONSchema     string
}

func (src *BaseDataSource) Name() string {
	return src.name
}

func (src *BaseDataSource) AuthScheme() AuthScheme {
	return src.authScheme
}

func (src *BaseDataSource) BindAirbyteSource(sourceDefinitionID, icon string) {
	src.AirbyteSourceDefinitionID = sourceDefinitionID
	src.Icon = icon
}

func (src *BaseDataSource) GetAirbyteConfig(config, secrets []byte) (interface{}, error) {
	return nil, errors.New(fmt.Sprintf("%s doesn't implement GetAirbyteConfig", src.name))
}

func (src *BaseDataSource) ToPb() *octopus.DataSource {
	authTypeInt := octopus.AuthType_value[string(src.authScheme)]
	return &octopus.DataSource{
		Name:                         src.Name(),
		AuthType:                     octopus.AuthType(authTypeInt),
		IconSvg:                      src.Icon,
		ConfigurationInputJSONSchema: src.ConfigInputJSONSchema,
		AirbyteSourceDefinitionId:    src.AirbyteSourceDefinitionID,
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
	return nil, errors.New(fmt.Sprintf("Data source %s ExchangeCode is not implemented", src.name))
}

func (src *OAuth2DataSource) ToPb() *octopus.DataSource {
	pbSource := src.BaseDataSource.ToPb()
	pbSource.AuthUrl = src.AuthURL
	return pbSource
}

type APIKeyDataSource struct {
	BaseDataSource
}
