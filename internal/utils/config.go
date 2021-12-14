package utils

import (
	"github.com/spf13/viper"
	"path/filepath"
)

func LoadYAMLConfig(cfgFilePath string) error {
	dir := filepath.Dir(cfgFilePath)
	file := filepath.Base(cfgFilePath)
	viper.AddConfigPath(dir)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}

func ReadYAMLConfig(cfgFilePath string, prefixKey string, cfg interface{}) error {
	if err := LoadYAMLConfig(cfgFilePath); err != nil {
		return err
	}

	if len(prefixKey) > 0 {
		if err := viper.UnmarshalKey(prefixKey, cfg); err != nil {
			return err
		}
	} else {
		if err := viper.Unmarshal(cfg); err != nil {
			return err
		}
	}
	return nil
}
