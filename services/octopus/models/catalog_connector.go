package models

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
