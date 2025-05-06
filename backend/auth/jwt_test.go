package auth

import (
	"os"
	"testing"
	"time"

	"github.com/yourorg/recipe-app/models"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Set test JWT secret
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	// Create a test user
	testUser := &models.User{
		ID:        1,
		Email:     "test@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Generate a token
	token, err := GenerateToken(testUser)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Token should not be empty
	if token == "" {
		t.Fatal("Generated token is empty")
	}

	// Validate token
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	// Verify claims
	if claims.UserID != testUser.ID {
		t.Errorf("UserID mismatch: got %d, want %d", claims.UserID, testUser.ID)
	}

	if claims.Email != testUser.Email {
		t.Errorf("Email mismatch: got %s, want %s", claims.Email, testUser.Email)
	}
}

func TestInvalidToken(t *testing.T) {
	// Set test JWT secret
	os.Setenv("JWT_SECRET", "test_secret")
	defer os.Unsetenv("JWT_SECRET")

	// Try to validate an invalid token
	_, err := ValidateToken("invalid.token.string")
	if err == nil {
		t.Fatal("Expected error for invalid token, but got nil")
	}
}