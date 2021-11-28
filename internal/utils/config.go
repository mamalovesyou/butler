package utils

import (
	"github.com/spf13/viper"
)

func LoadConfigWithViper(path, filename string) error {
	cfgFileName := FileNameWithoutExtension(filename)
	viper.AddConfigPath(path)
	viper.SetConfigName(cfgFileName)
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
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
