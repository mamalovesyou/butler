package cmd

import (
	"context"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/tools/victorinox"
	"github.com/spf13/cobra"
)

var (
	services   []string
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "SQL schema  migration cli for butler services",
		Long:  `SQL schema  migration cli for butler services`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			if len(args) <= 0 {
				logger.Fatal(ctx, "Missing goose command")
			}

			gooseCmd := args[0]
			if !victorinox.IsSupportedGooseCmd(gooseCmd) {
				logger.Fatalf(ctx, "Unsupported command: %s. Try using goose cli", gooseCmd)
			}

			arguments := []string{}
			if len(args) > 1 {
				arguments = append(arguments, args[1:]...)
			}

			// Load config
			cfg, err := victorinox.LoadConfig(configDir, configFileName)
			if err != nil {
				logger.Fatalf(ctx, "Failed to load config: %+v", err)
			}
			migrations := victorinox.NewGooseMigrations(cfg)

			for _, name := range services {
				migrations.RunGooseMigration(ctx, name, gooseCmd, args...)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&services, "services", "s", []string{"users"}, "Services name")
	rootCmd.AddCommand(migrateCmd)
}
