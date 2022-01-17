package connectors

type OAuthConnectorConfig struct {
	ClientID       string `envconfig:"client_id"`
	ClientSecret   string `envconfig:"client_secret"`
	DeveloperToken string `envconfig:"developer_token"`
}

type Config struct {
	Google      OAuthConnectorConfig `envconfig:"google"`
	Linkedin    OAuthConnectorConfig `envconfig:"linkedin"`
	RedirectURL string               `mapStructure:"redirectURL"`
}
