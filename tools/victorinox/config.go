package victorinox

import (
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/postgres"
	"github.com/butlerhq/butler/internal/utils"
	"github.com/spf13/viper"
)

type VictorioxConfig struct {
	Environment string
	Postgres    *postgres.PostgresConfig
	Logger      *logger.LoggerConfig
}

// LoadConfig return a VictorioxConfig
func LoadConfig(path, name string) (*VictorioxConfig, error) {
	if err := utils.LoadConfigWithViper(path, name); err != nil {
		return &VictorioxConfig{}, err
	}

	config := &VictorioxConfig{}
	if err := viper.UnmarshalKey("victorinox", config); err != nil {
		return &VictorioxConfig{}, err
	}

	return config, nil
}
