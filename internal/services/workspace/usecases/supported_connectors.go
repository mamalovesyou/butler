package usecases

import (
	"context"
	"github.com/matthieuberger/butler/internal/logger"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
)

var (
	GoogleConnectorConfig = oauth2.Config{
		ClientID:     "133098310007-oq0arc40c3o9821rmcq9oen5bncnn1ru.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/adwords"},
		Endpoint: google.Endpoint,
		RedirectURL: "http://localhost:3000/oauth",
	}

	LinkedInConnectorConfig = oauth2.Config{
		ClientID:     "86b45hlfzb80o7",
		ClientSecret: "mYjBhNuv81dAGWTk",
		Scopes:       []string{"r_liteprofile", "r_emailaddress"},
		Endpoint: linkedin.Endpoint,
		RedirectURL: "http://localhost:3000/oauth",
	}
)

func ExchangeOAuthCode(ctx context.Context, provider, code string) (*oauth2.Token, error) {
	switch provider {
	case "google":
		token, err := GoogleConnectorConfig.Exchange(ctx, code, oauth2.AccessTypeOffline)
		if err != nil {
			logger.Error(ctx, "Failed to exchange oauth google", zap.Error(err))
			return nil, ErrUnableToExchangeCode
		}
		return token, nil
	case "linkedin":
		token, err := LinkedInConnectorConfig.Exchange(ctx, code, oauth2.AccessTypeOffline)
		if err != nil {
			logger.Error(ctx, "Failed to exchange oauth linkedin", zap.Error(err))
			return nil, ErrUnableToExchangeCode
		}
		return token, nil
	default:
		return nil, ErrInvalidConnectorName
	}
}
