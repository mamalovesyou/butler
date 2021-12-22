package cmd

import (
	"github.com/spf13/cobra"
)

var (
	provisionCmd = &cobra.Command{
		Use:   "provision",
		Short: "Provision postgres postgres all services.",
		Long:  `Provision postgres postgres with credentials for all services.`,

		Run: func(cmd *cobra.Command, args []string) {
			//ctx := context.Background()

			// //Load config
			//cfg, err := victorinox.LoadConfig(configDir, configFileName)
			//if err != nil {
			//	logger.Fatalf(ctx, "Failed to load config: %+v", err)
			//}
			//logger.Info(ctx, "Loaded config", zap.Any("cfg", cfg))
			//registry := &victorinox.DBRegistry{
			//	RootDBConfig: cfg.RootPostgres,
			//	DBToProvision: []*postgres.PostgresConfig{cfg.UsersPostgres},
			//}
			//if err := registry.ProvisionAll(ctx); err != nil {
			//	logger.Fatalf(ctx, "Failed to provision protgres postgres: %v", err)
			//}
		},
	}
)

func init() {
	//rootCmd.AddCommand(provisionCmd)
}
