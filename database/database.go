package database

import (
	"fmt"
	"log"
	"plantlife/config"

	"github.com/jmoiron/sqlx"
)

func InitDB(config config.Config) (*sqlx.DB, error) {
	var err error

	connStr := fmt.Sprintf("port=%s host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DB_PORT, config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the database")

	return db, nil

}
