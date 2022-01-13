package connectors

type OAuthConnectorConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type ConnectorsConfig struct {
	Google   OAuthConnectorConfig
	Linkedin OAuthConnectorConfig
}
