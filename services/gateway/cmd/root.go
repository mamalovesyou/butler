package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/butlerhq/butler/internal/config"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/gateway"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	// Flags variables
	cfgFilePath string
	cfgKey      string

	gatewayCfg = gateway.DefaultGatewayConfig

	rootCmd = &cobra.Command{
		Use:   "gateway",
		Short: "HeyButler REST API Gateway",
		Long:  `HeyButler REST Gateway si a rest proxy server that proxy rest request to the right GRPC service.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			config.ReadConfig(cfgFilePath, cfgKey, &gatewayCfg)
			logger.Debug(ctx, "Starting users", zap.Any("config", gatewayCfg))
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
