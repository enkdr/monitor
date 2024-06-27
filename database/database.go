package database

import (
	"fmt"
	"time"

	"github.com/enkdr/monitor/config"
	"github.com/jmoiron/sqlx"
)

func InitDB(config config.Config) (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB

	connStr := fmt.Sprintf("port=%s host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DB_PORT, config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	// Retry with exponential backoff
	maxAttempts := 5 // Maximum number of retry attempts
	attempt := 0
	for attempt < maxAttempts {
		db, err = sqlx.Connect("postgres", connStr)
		if err != nil {
			// Retry if there's an error connecting
			fmt.Printf("Attempt %d: Failed to connect to database: %v\n", attempt+1, err)
			attempt++
			backoff := time.Duration(1<<uint(attempt)) * time.Second // Exponential backoff
			fmt.Printf("Retrying in %v...\n", backoff)
			time.Sleep(backoff)
		} else {
			break // Connected successfully
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxAttempts, err)
	}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil

}
