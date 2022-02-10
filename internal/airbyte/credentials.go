package airbyte

import "golang.org/x/oauth2"

const OAUTH_METHOD = "oAuth2.0"

type Credentials struct {
	AuthMethod   string `json:"auth_method"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

func NewCredentialsFromOAuth(cfg *oauth2.Config, token *oauth2.Token) *Credentials {
	return &Credentials{
		AuthMethod:   OAUTH_METHOD,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RefreshToken: token.RefreshToken,
	}
}
