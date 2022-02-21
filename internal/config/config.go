package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v6"
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

func ReadConfig(cfgFilePath string, prefixKey string, cfg interface{}) error {
	fmt.Println(os.Environ())
	LoadYAMLConfig(cfgFilePath)

	if len(prefixKey) > 0 {
		if err := viper.UnmarshalKey(prefixKey, cfg); err != nil {
			return err
		}
	} else {
		if err := viper.Unmarshal(cfg); err != nil {
			return err
		}
	}

	return env.Parse(cfg)
}
