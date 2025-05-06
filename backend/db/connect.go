package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	// Import pgx driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/yourorg/recipe-app/config"
)

// Connect creates a database connection using environment variables
func Connect() (*sql.DB, error) {
	// Get database connection parameters from environment
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASS", "postgres")
	dbname := config.GetEnv("DB_NAME", "recipe_app")

	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to database
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test connection
	if err := db.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Configure connection pool
	maxOpenConns, err := strconv.Atoi(config.GetEnv("DB_MAX_OPEN_CONNS", "25"))
	if err == nil {
		db.SetMaxOpenConns(maxOpenConns)
	}

	maxIdleConns, err := strconv.Atoi(config.GetEnv("DB_MAX_IDLE_CONNS", "25"))
	if err == nil {
		db.SetMaxIdleConns(maxIdleConns)
	}
	
	return db, nil
}