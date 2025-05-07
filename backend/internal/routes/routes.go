package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/TheAlonso95/recipe-app/internal/handlers"
)

// SetupRoutes configures all API routes
func SetupRoutes(db *sql.DB) (http.Handler, error) {
	// Check if the database connection is valid
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection error: %v", err)
	}

	// Create a new router
	mux := http.NewServeMux()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)

	// Authentication routes
	mux.HandleFunc("/api/auth/register", authHandler.Register)
	mux.HandleFunc("/api/auth/login", authHandler.Login)

	return mux, nil
}
