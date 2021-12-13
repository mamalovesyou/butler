package victorinox

import (
	"github.com/butlerhq/butler/internal/postgres"
)

type DBRegistry struct {
	RootDBConfig  *postgres.PostgresConfig
	DBToProvision []*postgres.PostgresConfig
}

//func (regitry *DBRegistry) ProvisionAll(ctx context.Context) error {
//	for _, config := range regitry.DBToProvision {
//		logger.Infof(ctx, "Provisioning Postgres DB %v", config)
//		if err := provisionDB(ctx, regitry.RootDBConfig, config); err != nil {
//			logger.Fatalf(ctx, "❌ Unable to provision postgres db %s: %v", config.Name, err)
//			return err
//		}
//	}
//	return nil
//}

//func provisionDB (ctx context.Context, rootConfig *postgres.PostgresConfig, provisionCfg *postgres.PostgresConfig) error {
//	logger.Infof(ctx, "Provisioning database: %s", provisionCfg.Name)
//
//	// Attempt to connect db with root user
//	db := postgres.NewPostgres(provisionCfg)
//	if err := db.ConnectLoop(5*time.Second, true); err != nil {
//		logger.Error(ctx, "❌ Failed to connect database", zap.Error(err))
//		logger.Info(ctx, "Trying to connect using root config")
//		db = postgres.NewPostgres(rootConfig)
//		if err := db.ConnectLoop(5*time.Second, true); err != nil {
//			logger.Error(ctx, "❌ Failed to connect database", zap.Error(err))
//			return err
//		}
//	}
//	logger.Info(ctx, "✅ Database connected")
//
//	// Create user and grants privileges if not exists
//	if err := db.CreateUserAndGrantsPrivilege(); err != nil {
//		logger.Fatalf(ctx, "❌ Failed to create user %s for database %s: %v", provisionCfg.User, provisionCfg.Name, err)
//	}
//	logger.Infof(ctx, "✅ User %s for database %s", provisionCfg.User, provisionCfg.Name)
//
//	// Creating database if not exists
//	if err := db.CreateDatabaseIfNotExists(); err != nil {
//		logger.Fatal(ctx, "❌ Failed to create database", zap.Error(err), zap.String("db", provisionCfg.Name))
//	}
//	logger.Infof(ctx, "✅ Database: %s", provisionCfg.Name)
//
//	// Create uuid extension if not exists
//	if err := db.CreateUUIDExtension(); err != nil {
//		logger.Fatal(ctx, "❌ Failed to create UUID extension", zap.Error(err))
//	}
//	logger.Info(ctx, "✅ UUID extension")
//
//	return nil
//}
