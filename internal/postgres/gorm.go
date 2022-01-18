package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultConnectTimeout = 5 * time.Second
)

type PostgresGorm struct {
	Config *Config
	DB     *gorm.DB
	SqlDB  *sql.DB
}

// NewPostgresGorm create a new gorm wrapper
func NewPostgresGorm(cfg *Config) *PostgresGorm {
	return &PostgresGorm{Config: cfg}
}

// Connect tries to connect to the DB under given DSN
func (pg *PostgresGorm) Connect() (err error) {
	dsn := pg.Config.GetConnectionURI()
	if pg.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}
	if pg.SqlDB, err = pg.DB.DB(); err != nil {
		return err
	}
	return nil
}

// ConnectLoop tries to connect to the DB under given DSN
// in a loop until connection succeeds. timeout specifies the timeout for the
// loop
func (pg *PostgresGorm) ConnectLoop(timeout time.Duration) error {
	dsn := pg.Config.GetConnectionURI()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return fmt.Errorf("Failed to connect DB after %s timeout", timeout)

		case <-ticker.C:
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				pg.DB = db
				pg.SqlDB, err = pg.DB.DB()
				if err != nil {
					return err
				}
				return nil
			}
		}
	}
}

// Close database connection
func (pg *PostgresGorm) Close() error {
	err := pg.SqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}

// Tables return a list of all tables in public schemas
func (pg *PostgresGorm) Tables() []string {
	var tables []string
	if err := pg.DB.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		panic(err)
	}
	return tables
}

// FlushAll will delete all records for all available tables. Useful for end to end tests
// when we want to reset data in database
func (pg *PostgresGorm) TruncateTables(tables []string) error {
	return pg.DB.Transaction(func(tx *gorm.DB) error {
		for _, t := range tables {
			sql := fmt.Sprintf(`TRUNCATE TABLE %s RESTART IDENTITY CASCADE;`, t)
			if err := tx.Exec(sql).Error; err != nil {
				return fmt.Errorf("Failed to truncate table %s > Error: %v", t, err)
			}
		}
		return nil
	})
}

// FlushAll will delete all records for all available tables. Useful for end to end tests
// when we want to reset data in database
func (pg *PostgresGorm) FlushAll() {
	tables := pg.Tables()
	if err := pg.TruncateTables(tables); err != nil {
		panic(err)
	}
}
