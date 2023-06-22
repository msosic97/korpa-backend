package db

// All configuration for connection on database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/msosic97/korpa-backend/config"
	_ "github.com/go-sql-driver/mysql"
)




var DB *sql.DB

func validateDBConfig(config *config.Config) error {
	if len(config.DBName) == 0 {
		return fmt.Errorf("no database name provided")
	}
	if len(config.DBUser) == 0 {
		return fmt.Errorf("no user name provided")
	}
	if len(config.DBHost) == 0 {
		return fmt.Errorf("no host provided")
	}
	if len(config.DBPassword) == 0 {
		return fmt.Errorf("no password provided")
	}
	if len(config.DBPort) == 0 {
		return fmt.Errorf("no port provided")
	}
	return nil
}

// Connection to DB
func StartDB(config *config.Config) (*sql.DB, error) {
	if err := validateDBConfig(config); err != nil {
		return nil, err
	}

	// DSN is a string that contains the necessary information to establish a connection to a specific data source or database.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Successfully connected to the database!")

	DB = db
	return db, nil
}