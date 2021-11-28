package workspace

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	// Flags variables
	configFileName string
	configDir      string

	rootCmd = &cobra.Command{
		Use:   "workspace",
		Short: "Workspace is a service that provide organizations and workspaces management.",
		Long:  `Workspace is a service that provide organizations and workspaces management."`,
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
	rootCmd.PersistentFlags().StringVarP(&configFileName, "filename", "f", "config.yml", "Config file name")
	rootCmd.PersistentFlags().StringVarP(&configDir, "dir", "", ".", "Config file directory")
}
