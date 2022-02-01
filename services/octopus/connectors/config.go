package connectors

type OAuthConnectorConfig struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
}

type Config struct {
	Google      OAuthConnectorConfig `envPrefix:"GOOGLE_"`
	Linkedin    OAuthConnectorConfig `envPrefix:"LINKEDIN_"`
	RedirectURL string               `mapStructure:"redirectURL"`
}
