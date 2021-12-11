package config

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/redis"
	"github.com/spf13/viper"
	"log"
)

// App config struct
type Config struct {
	Server   ServerConfig
	Postgres *postgres.PostgresConfig
	Redis    *redis.RedisConfig
	Jaeger   *logger.JaegerConfig
	Logger   *logger.LoggerConfig
}

// Server config struct
type ServerConfig struct {
	PlatformEnv string
	Port        string
	//PprofPort         string
	//Mode              string
	//JwtSecretKey      string
	//CookieName        string
	//ReadTimeout       time.Duration
	//WriteTimeout      time.Duration
	//SSL               bool
	//CtxDefaultTimeout time.Duration
	//CSRF              bool
	//Debug             bool
	//MaxConnectionIdle time.Duration
	//Timeout           time.Duration
	//MaxConnectionAge  time.Duration
	//Time              time.Duration
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, ErrConfigNotFound
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

// Get config
func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
