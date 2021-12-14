package postgres

import (
	"database/sql"
	"fmt"
	"time"
)

type Postgres struct {
	Config *PostgresConfig
	DB     *sql.DB
}

func NewPostgres(cfg *PostgresConfig) *Postgres {
	return &Postgres{Config: cfg}
}

// Connect tries to connect to the DB under given DSN
func (pg *Postgres) Connect(cfg *PostgresConfig, withoutDB bool) (err error) {
	var uri string
	if uri = cfg.GetConnectionURI(); withoutDB {
		uri = cfg.GetConnectionURIWithoutDB()
	}
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	pg.DB = db
	return
}

// ConnectLoop tries to connect to the DB under given DSN
// in a loop until connection succeeds. timeout specifies the timeout for the
// loop
func (pg *Postgres) ConnectLoop(timeout time.Duration, withoutDB bool) (err error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			err = fmt.Errorf("Failed to connect DB after %s timeout", timeout)
			return err

		case <-ticker.C:
			if err = pg.Connect(pg.Config, withoutDB); err == nil {
				return
			}
		}
	}
}

// Close database connection
func (pg *Postgres) Close() {
	pg.DB.Close()
}

func (pg *Postgres) CreateUUIDExtension() (err error) {
	if _, err = pg.DB.Exec(
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`); err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) CreateDatabaseIfNotExists() error {
	rows, err := pg.DB.Query(`SELECT 1 AS result FROM pg_database WHERE datname=$1;`, pg.Config.Name)
	if err != nil {
		return err
	}
	// No datbase found
	if !rows.Next() {
		// All query statements don't support parameters in postgres.
		// We an't do DB.Exec(`CREATE DATABASE $1 ENCODING='UTF8';`, pg.Config.Name)
		query := fmt.Sprintf("CREATE DATABASE %s ENCODING='UTF8';", pg.Config.Name)
		if _, err = pg.DB.Exec(query); err != nil {
			return err
		}
	}
	return nil
}

func (pg *Postgres) CreateUserAndGrantsPrivilege() error {
	rows, err := pg.DB.Query(`SELECT 1 AS result FROM pg_roles WHERE rolname=$1`, pg.Config.User)
	if err != nil {
		return err
	}
	// No user found with username
	if !rows.Next() {
		query := `CREATE USER $1 WITH PASSWORD $2; GRANT ALL PRIVILEGES ON DATABASE $3 to user1;`
		if _, err = pg.DB.Exec(query, pg.Config.User, pg.Config.Password, pg.Config.Name); err != nil {
			return err
		}
	}
	return nil
}
