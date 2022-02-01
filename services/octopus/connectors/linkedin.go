package connectors

import (
	"context"
	"fmt"

	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"

	linkedin_ads "github.com/butlerhq/butler/integrations/linkedin-ads"

	"github.com/butlerhq/butler/api/services/octopus/v1"

	"github.com/butlerhq/butler/services/octopus/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

var LINKEDIN_NAME = "LINKEDIN"

type LinkedinConnector struct {
	config oauth2.Config
}

func NewLinkedinConnector(cfg OAuthConnectorConfig, redirectURL string) *LinkedinConnector {
	return &LinkedinConnector{
		config: oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Scopes:       []string{"r_liteprofile", "r_emailaddress", "rw_ads"},
			Endpoint:     linkedin.Endpoint,
			RedirectURL:  redirectURL,
		},
	}
}

func (lc *LinkedinConnector) Name() string {
	return LINKEDIN_NAME
}

func (lc *LinkedinConnector) AuthScheme() models.AuthScheme {
	return models.OAUTH2
}

func (lc *LinkedinConnector) SVGIcon() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="96px" height="96px"><path fill="#0288D1" d="M42,37c0,2.762-2.238,5-5,5H11c-2.761,0-5-2.238-5-5V11c0-2.762,2.239-5,5-5h26c2.762,0,5,2.238,5,5V37z"/><path fill="#FFF" d="M12 19H17V36H12zM14.485 17h-.028C12.965 17 12 15.888 12 14.499 12 13.08 12.995 12 14.514 12c1.521 0 2.458 1.08 2.486 2.499C17 15.887 16.035 17 14.485 17zM36 36h-5v-9.099c0-2.198-1.225-3.698-3.192-3.698-1.501 0-2.313 1.012-2.707 1.99C24.957 25.543 25 26.511 25 27v9h-5V19h5v2.616C25.721 20.5 26.85 19 29.738 19c3.578 0 6.261 2.25 6.261 7.274L36 36 36 36z"/></svg>`
}

func (lc *LinkedinConnector) AuthURL() string {
	return lc.config.AuthCodeURL("", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func (lc *LinkedinConnector) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := lc.config.Exchange(ctx, code, oauth2.AccessTypeOffline)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (lc *LinkedinConnector) ToPb() *octopus.CatalogConnector {
	authTypeInt := octopus.AuthType_value[string(lc.AuthScheme())]
	return &octopus.CatalogConnector{
		Name:     lc.Name(),
		AuthUrl:  lc.AuthURL(),
		AuthType: octopus.AuthType(authTypeInt),
		IconSvg:  lc.SVGIcon(),
	}
}

func (lc *LinkedinConnector) ListAccounts(ctx context.Context, secrets *models.ConnectorSecrets) (*octopus.ListAccountsResponse, error) {
	linkedin_client := linkedin_ads.NewLinkedinClient(secrets.AccessToken)
	resp, err := linkedin_client.ListAccounts()
	if err != nil {
		logger.Error(ctx, "Failed to list linkedin accounts", zap.Error(err))
		return &octopus.ListAccountsResponse{}, nil
	}

	result := make([]*octopus.ProviderAccount, len(resp.Accounts))
	for i, acc := range resp.Accounts {
		result[i] = &octopus.ProviderAccount{
			Name: acc.Name,
			Id:   fmt.Sprintf("%d", acc.Id),
			Test: acc.Test,
		}
	}

	return &octopus.ListAccountsResponse{
		Accounts: result,
	}, nil

}
