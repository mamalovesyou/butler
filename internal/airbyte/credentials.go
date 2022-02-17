package airbyte

import "golang.org/x/oauth2"

const OAUTH_METHOD = "oAuth2.0"

type Credentials struct {
	AuthMethod     string `json:"auth_method,omitempty"`
	ClientID       string `json:"client_id,omitempty"`
	ClientSecret   string `json:"client_secret,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
	RefreshToken   string `json:"refresh_token,omitempty"`
	DeveloperToken string `json:"developer_token,omitempty"`
}

func NewCredentialsFromOAuth(cfg *oauth2.Config, token *oauth2.Token) *Credentials {
	return &Credentials{
		AuthMethod:   OAUTH_METHOD,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RefreshToken: token.RefreshToken,
		AccessToken:  token.AccessToken,
	}
}
