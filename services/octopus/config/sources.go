package config

type OAuthSourceConfig struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
}

type SourcesConfig struct {
	Google      OAuthSourceConfig `envPrefix:"GOOGLE_"`
	Linkedin    OAuthSourceConfig `envPrefix:"LINKEDIN_"`
	RedirectURL string            `mapStructure:"redirectURL"`
}
