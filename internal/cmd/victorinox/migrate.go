package victorinox

import (
	"context"
	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Apply or rollback SQL migrations for given service.",
		Long:  `Apply or rollback SQL migrations for given service.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			cfgFileName := utils.FileNameWithoutExtension(configFileName)

			viper.AddConfigPath(configDir)
			viper.SetConfigName(cfgFileName)
			viper.SetConfigType("yaml")

			err := viper.ReadInConfig()
			if err != nil {
				logger.Error(ctx, "❌ Failed to read config", zap.Error(err))
			}

			//for _, prefix := range services.ServiceWithPostgresDBList {
			//
			//	// Verify prefix key is present in config
			//	if ok := viper.InConfig(prefix); !ok {
			//		// TODO: Move error to dedicated file
			//		err := errors.New(fmt.Sprintf("Key: %s is missing", prefix))
			//		logger.Fatal(ctx, "❌ Unable to load config. ", zap.Error(err))
			//	}
			//
			//	files, ok := registry.MigrationsRegistry[prefix]
			//	if !ok {
			//		err := errors.New(fmt.Sprintf("No migrations found for service: %s", prefix))
			//		logger.Fatal(ctx, "❌ Unable to find migration files. ", zap.Error(err))
			//	}
			//	fs := http.FS(files)
			//	source, err := httpfs.New(fs, "migrations")
			//	if err != nil {
			//		logger.Fatal(ctx, "❌ Failed to load migrations.", zap.Error(err))
			//	}
			//	logger.Info(ctx, "Migrate database schema...")
			//	config := &postgres.PostgresConfig{}
			//	key := fmt.Sprintf("%s.postgres", prefix)
			//	if err := viper.UnmarshalKey(key, config); err != nil {
			//		logger.Fatal(ctx, "❌ Failed to unmarshal config.", zap.Error(err))
			//	}
			//	logger.Info(ctx, "Got", zap.String("key", key), zap.Any("config", config))
			//
			//	postgresInstance := postgres.NewPostgres(config)
			//	if err := postgresInstance.ConnectLoop(5*time.Second, false); err != nil {
			//		logger.Fatal(ctx, "❌ Failed to connect database", zap.Error(err))
			//	}
			//	logger.Info(ctx, "✅ Successfully connected to postgres")
			//	driver, err := migrate_postgres.WithInstance(postgresInstance.DB, &migrate_postgres.Config{})
			//	m, err := migrate.NewWithInstance("httpfs", source, config.Name, driver)
			//	if err != nil {
			//		logger.Fatal(ctx, "❌ Failed to initialize migrate instance.", zap.Error(err))
			//	}
			//	if err := m.Up(); err != nil {
			//		if err == migrate.ErrNoChange {
			//			logger.Infof(ctx, "%s - No change detected.", prefix)
			//		} else {
			//			logger.Fatal(ctx, "❌ Failed to apply migrations.", zap.Error(err))
			//		}
			//	}
			//	logger.Infof(ctx, "✅ %s - Successfully applied migrations", prefix)
			//}
		},
	}
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}
