package db

// All configuration for connection on database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/msosic97/korpa-backend/config"
	_ "github.com/go-sql-driver/mysql"
)


// DSN is a string that contains the necessary information to establish a connection to a specific data source or database.

var DB *sql.DB

//  Connection to DB
func StartDB(config *config.Config) (*sql.DB, error) {
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