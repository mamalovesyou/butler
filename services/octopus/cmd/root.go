package cmd

import (
	"context"
	"fmt"
	"os"

	config2 "github.com/butlerhq/butler/services/octopus/config"

	"github.com/butlerhq/butler/internal/config"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var (
	// Flags variables
	cfgFilePath string
	cfgKey      string

	cfgService = config2.DefaultServiceConfig

	rootCmd = &cobra.Command{
		Use:   "butler-octopus",
		Short: "butler-octopus is a service that provide octopus credential management.",
		Long:  `butler-octopus is a service that provide octopus credential management.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			config.ReadConfig(cfgFilePath, cfgKey, &cfgService)
			logger.Debug(ctx, "Starting octopus", zap.Any("config", cfgService))
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
