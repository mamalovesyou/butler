package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/butlerhq/butler/internal/config"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/users"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var (
	// Flags variables
	cfgFilePath string
	cfgKey      string

	usersConfig = users.DefaultServiceConfig

	rootCmd = &cobra.Command{
		Use:   "butler-users",
		Short: "butler-users is a service that provide user, permissions, workspace and team management.",
		Long:  `butler-users is a service that provide user, permissions, workspace and team management.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			config.ReadConfig(cfgFilePath, cfgKey, &usersConfig)
			logger.Debug(ctx, "Starting users", zap.Any("config", usersConfig))
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
