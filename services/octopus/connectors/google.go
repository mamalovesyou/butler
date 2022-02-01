package connectors

import (
	"context"
	"fmt"

	"github.com/butlerhq/butler/internal/logger"
	"go.uber.org/zap"

	"github.com/butlerhq/google-ads-go/ads"
	"github.com/butlerhq/google-ads-go/services"

	"github.com/butlerhq/butler/api/services/octopus/v1"

	"github.com/butlerhq/butler/services/octopus/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GOOGLE_NAME = "GOOGLE"

type GoogleConnector struct {
	config oauth2.Config
}

func NewGoogleConnector(cfg OAuthConnectorConfig, redirectURL string) *GoogleConnector {
	return &GoogleConnector{
		config: oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Scopes: []string{
				"https://www.googleapis.com/auth/user.organization.read",
				"https://www.googleapis.com/auth/adwords",
			},
			Endpoint:    google.Endpoint,
			RedirectURL: redirectURL,
		},
	}
}

func (gc *GoogleConnector) Name() string {
	return GOOGLE_NAME
}

func (gc *GoogleConnector) AuthScheme() models.AuthScheme {
	return models.OAUTH2
}

func (gc *GoogleConnector) SVGIcon() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="144px" height="144px"><path fill="#FFC107" d="M43.611,20.083H42V20H24v8h11.303c-1.649,4.657-6.08,8-11.303,8c-6.627,0-12-5.373-12-12c0-6.627,5.373-12,12-12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C12.955,4,4,12.955,4,24c0,11.045,8.955,20,20,20c11.045,0,20-8.955,20-20C44,22.659,43.862,21.35,43.611,20.083z"/><path fill="#FF3D00" d="M6.306,14.691l6.571,4.819C14.655,15.108,18.961,12,24,12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C16.318,4,9.656,8.337,6.306,14.691z"/><path fill="#4CAF50" d="M24,44c5.166,0,9.86-1.977,13.409-5.192l-6.19-5.238C29.211,35.091,26.715,36,24,36c-5.202,0-9.619-3.317-11.283-7.946l-6.522,5.025C9.505,39.556,16.227,44,24,44z"/><path fill="#1976D2" d="M43.611,20.083H42V20H24v8h11.303c-0.792,2.237-2.231,4.166-4.087,5.571c0.001-0.001,0.002-0.001,0.003-0.002l6.19,5.238C36.971,39.205,44,34,44,24C44,22.659,43.862,21.35,43.611,20.083z"/></svg>`
}

func (gc *GoogleConnector) AuthURL() string {
	return gc.config.AuthCodeURL("", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func (gc *GoogleConnector) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := gc.config.Exchange(ctx, code, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (gc *GoogleConnector) ListAccounts(ctx context.Context, secrets *models.ConnectorSecrets) (*octopus.ListAccountsResponse, error) {

	client, err := ads.NewClient(&ads.GoogleAdsClientParams{
		ClientID:     gc.config.ClientID,
		ClientSecret: gc.config.ClientSecret,
		RefreshToken: secrets.RefreshToken,
	})

	if err != nil {
		panic(err)
	}

	customerService := services.NewCustomerServiceClient(client.Conn())
	response, err := customerService.ListAccessibleCustomers(client.Context(), &services.ListAccessibleCustomersRequest{})
	if err != nil {
		logger.Error(ctx, "Unable to list accessible customers", zap.Error(err), zap.String("provider", gc.Name()))
		return &octopus.ListAccountsResponse{}, err
	}

	result := make([]*octopus.ProviderAccount, len(response.ResourceNames))
	for i, name := range response.ResourceNames {
		customer, err := customerService.GetCustomer(client.Context(), &services.GetCustomerRequest{
			ResourceName: name,
		})

		if err != nil {
			logger.Error(ctx, "Unable to retrieve account", zap.Error(err), zap.String("provider", gc.Name()))
			continue
		}

		result[i] = &octopus.ProviderAccount{
			Name:     *customer.DescriptiveName,
			Id:       fmt.Sprintf("%d", customer.Id),
			Test:     *customer.TestAccount,
			Currency: *customer.CurrencyCode,
		}
	}

	return &octopus.ListAccountsResponse{
		Accounts: result,
	}, nil
}

func (gc *GoogleConnector) ToPb() *octopus.CatalogConnector {
	authTypeInt := octopus.AuthType_value[string(gc.AuthScheme())]
	return &octopus.CatalogConnector{
		Name:     gc.Name(),
		AuthUrl:  gc.AuthURL(),
		AuthType: octopus.AuthType(authTypeInt),
		IconSvg:  gc.SVGIcon(),
	}
}
