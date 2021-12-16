package postgres

import (
	"fmt"
)

var (
	DefaultPostgresConfig = PostgresConfig{
		Host:     "postgres",
		Port:     "5432",
		Name:     "postgres",
		User:     "postgres",
		Password: "password",
	}
)

// PostgresConfig contains infos needed to open a postgres connection
type PostgresConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

// GetConnectionURI build and return database connection URI
// host=postgres port=5432 user=postgres_user dbname=postgres_name  password=strongPassword
func (cfg *PostgresConfig) GetConnectionURI() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)
}

// GetConnectionURL build and return database connection URL
// ex. postgres://username:password@localhost:5432/database_name
func (cfg *PostgresConfig) GetConnectionURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}

// GetConnectionURI build and return database connection URI
// host=postgres port=5432 user=postgres_user dbname=postgres_name  password=strongPassword
func (cfg *PostgresConfig) GetConnectionURIWithoutDB() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s sslmode=disable password=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password)
}

// GetConnectionURL build and return database connection URL
// ex. postgres://username:password@localhost:5432/database_name
func (cfg *PostgresConfig) GetConnectionURLWithoutDB() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/",
		cfg.User, cfg.Password, cfg.Host, cfg.Port)
}
