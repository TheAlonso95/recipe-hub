package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/yourorg/recipe-app/config"
	"github.com/yourorg/recipe-app/db"
	"github.com/yourorg/recipe-app/internal/handlers"
	"github.com/yourorg/recipe-app/internal/repositories"
)

func setupRoutes(r *mux.Router, authHandler *handlers.AuthHandler) {
	// API routes
	api := r.PathPrefix("/api").Subrouter()
	
	// Auth routes
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")
}

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	log.Println("Connected to PostgreSQL database successfully!")

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userRepo)

	// Create router
	router := mux.NewRouter()
	
	// Set up routes
	setupRoutes(router, authHandler)

	// Create server
	port := config.GetEnv("PORT", "8080")
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine so it doesn't block
	go func() {
		log.Printf("Server starting on port %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
	log.Println("Server gracefully stopped")
}