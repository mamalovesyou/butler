package config

import (
	"path/filepath"

	"github.com/butlerhq/butler/internal/utils"
	"github.com/spf13/viper"
)

const flagTagName = "flag"

func LoadYAMLConfig(cfgFilePath string) error {
	dir := filepath.Dir(cfgFilePath)
	file := filepath.Base(cfgFilePath)
	viper.AddConfigPath(dir)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}

func LoadEnvConfig() {
	for k, v := range utils.GetEnvironMap() {
		viperKey := utils.EnvToViperKey(k)
		viper.Set(viperKey, v)
	}
}

func ReadConfig(cfgFilePath string, prefixKey string, cfg interface{}) error {
	LoadYAMLConfig(cfgFilePath)
	LoadEnvConfig()

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
