package models

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/butlerhq/butler/api/services/octopus/v1"
)

type CatalogConnector interface {
	SVGIcon() string
	Name() string
	AuthScheme() AuthScheme
	AuthURL() string
	ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error)
	ToPb() *octopus.CatalogConnector
}
