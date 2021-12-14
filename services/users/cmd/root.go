package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var (
	// Flags variables
	cfgFilePath string
	cfgKey      string

	rootCmd = &cobra.Command{
		Use:   "butler-users",
		Short: "butler-users is a service that provide user, permissions, workspace and team management.",
		Long:  `butler-users is a service that provide user, permissions, workspace and team management.`,
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
