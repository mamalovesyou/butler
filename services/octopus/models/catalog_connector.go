package models

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/butlerhq/butler/api/services/octopus/v1"
)

type ConnectorSecrets struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken string `json:"access_token,omitempty"`
	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string `json:"refresh_token,omitempty"`
	// DeveloperToken is a token that authorizes and authenticates
	// the requests.
	DeveloperToken string `json:"developer_token,omitempty"`
	// DeveloperToken is a token that authorizes and authenticates
	// the requests.
	APIKey string `json:"api_key,omitempty"`
}

type CatalogConnector interface {
	SVGIcon() string
	Name() string
	AuthScheme() AuthScheme
	AuthURL() string
	ExchangeCode(context.Context, string) (*oauth2.Token, error)
	ListAccounts(context.Context, *ConnectorSecrets) (*octopus.ListAccountsResponse, error)
	ToPb() *octopus.CatalogConnector
}
