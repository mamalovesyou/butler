package cmd

import (
	"context"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/tools/victorinox"
	"github.com/spf13/cobra"
)

var (
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

			migrations.RunGooseMigration(ctx, service, gooseCmd, args...)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&service, "service", "s", "users", "Services name")
	rootCmd.PersistentFlags().StringVarP(&victorinoxCfg.Postgres.Host, "pg-host", "", victorinoxCfg.Postgres.Host, "Postgres host")
	rootCmd.PersistentFlags().StringVarP(&victorinoxCfg.Postgres.Port, "pg-port", "", victorinoxCfg.Postgres.Port, "Postgres port")
	rootCmd.PersistentFlags().StringVarP(&victorinoxCfg.Postgres.Name, "pg-name", "", victorinoxCfg.Postgres.Name, "Postgres database name")
	rootCmd.PersistentFlags().StringVarP(&victorinoxCfg.Postgres.User, "pg-user", "", victorinoxCfg.Postgres.User, "Postgres user")
	rootCmd.PersistentFlags().StringVarP(&victorinoxCfg.Postgres.Password, "pg-password", "", victorinoxCfg.Postgres.Password, "Postgres password")
	rootCmd.AddCommand(migrateCmd)
}
