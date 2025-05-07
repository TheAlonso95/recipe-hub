package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/TheAlonso95/recipe-app/internal/auth"
)

// AuthMiddleware checks if the request has a valid JWT token
// KeyUserID is the key used to store the user ID in the request context
const KeyUserID = "user_id"

// AuthMiddleware checks if the request has a valid JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check if the header has the Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Authorization header must be in format: Bearer {token}", http.StatusUnauthorized)
			return
		}

		// Validate the token
		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add user ID to request context
		ctx := context.WithValue(r.Context(), KeyUserID, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
