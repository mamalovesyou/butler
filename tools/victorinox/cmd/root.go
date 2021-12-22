package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/butlerhq/butler/internal/config"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/tools/victorinox"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type TestStruct struct {
	Environment string `env="ENVIRONMENT"`
}

// rootCmd represents the base command when called without any subcommands
var (
	// Flags variables
	cfgFilePath string
	cfgKey      string

	victorinoxCfg = victorinox.DefaultVictorinoxConfig

	rootCmd = &cobra.Command{
		Use:   "victorinox",
		Short: "Victorinox is a collection of tool to help manage infrastructure and postgres.",
		Long:  `Victorinox is a collection of tool to help manage infrastructure and postgres.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			config.ReadConfig(cfgFilePath, cfgKey, &victorinoxCfg)
			logger.Debug(ctx, "Starting victorinox", zap.Any("config", victorinoxCfg))
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c", "config.yml", "Config file path")
	rootCmd.PersistentFlags().StringVarP(&cfgKey, "key", "", "", "Config key")
}
