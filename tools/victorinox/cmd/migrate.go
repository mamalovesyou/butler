package cmd

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/tools/victorinox"
	"github.com/spf13/cobra"
)

var (
	allService bool
	service    string
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

			migrations := victorinox.NewGooseMigrations(&victorinoxCfg)
			if allService {
				migrations.RunGooseMigrationForAllServices(ctx, gooseCmd, args...)
			} else {
				migrations.RunGooseMigrationForService(ctx, service, gooseCmd, args...)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&service, "service", "s", "users", "Services name")
	rootCmd.PersistentFlags().BoolVarP(&allService, "all", "", false, "Migrate all services")
	rootCmd.AddCommand(migrateCmd)
}
