package routes

import (
	"database/sql"
	"net/http"

	"github.com/yourorg/recipe-app/handlers"
)

// SetupRoutes configures all API routes
func SetupRoutes(db *sql.DB) http.Handler {
	// Create a new router
	mux := http.NewServeMux()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)

	// Authentication routes
	mux.HandleFunc("POST /api/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/auth/login", authHandler.Login)

	return mux
}