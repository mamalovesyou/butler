package victorinox

import (
	"context"
	"github.com/butlerhq/butler/internal/utils"
	"go.uber.org/zap"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	provisionCmd = &cobra.Command{
		Use:   "provision",
		Short: "Provision new databases with credentials for all services.",
		Long:  `Provision new databases with credentials for all services.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			cfgFileName := utils.FileNameWithoutExtension(configFileName)

			viper.AddConfigPath(configDir)
			viper.SetConfigName(cfgFileName)
			viper.SetConfigType("yaml")

			err := viper.ReadInConfig()
			if err != nil {
				logger.Fatal(ctx, "❌ Failed to read config", zap.Error(err))
			}

			//for _, prefix := range services.ServiceWithPostgresDBList {
			//	logger.Infof(ctx, "Provisioning database for service %s", prefix)
			//	config := &postgres.PostgresConfig{}
			//	key := fmt.Sprintf("%s.postgres", prefix)
			//
			//	if err := viper.UnmarshalKey(key, config); err != nil {
			//		logger.Fatal(ctx, "❌ Failed to unmarshal config", zap.Error(err))
			//	}
			//	db := postgres.NewPostgres(config)
			//	if err := db.ConnectLoop(5*time.Second, true); err != nil {
			//		logger.Fatal(ctx, "❌ Failed to connect database", zap.Error(err))
			//	}
			//
			//	logger.Info(ctx, "✅ Successfully connected to postgres")
			//
			//	// Creating database
			//	if err = db.CreateDatabaseIfNotExists(); err != nil {
			//		logger.Fatal(ctx, "❌ Failed to create database", zap.Error(err), zap.String("db", config.Name))
			//	}
			//	logger.Infof(ctx, "✅ Successfully created database: %s", config.Name)
			//
			//	// Create uuid extension
			//	if err = db.CreateUUIDExtension(); err != nil {
			//		logger.Fatal(ctx, "❌ Failed to create UUID extension", zap.Error(err))
			//	}
			//	logger.Info(ctx, "✅ Created UUID extension")
			//
			//	// Create user and grants privilledge
			//	if err = db.CreateUserAndGrantsPrivilege(); err != nil {
			//		logger.Fatalf(ctx, "❌ Failed to create user %s for database %s", config.User, config.Name, err)
			//	}
			//	logger.Infof(ctx, "✅ Successfully created user %s for database %s", config.User, config.Name)
			//
			//}
		},
	}
)

func init() {
	rootCmd.AddCommand(provisionCmd)
}
