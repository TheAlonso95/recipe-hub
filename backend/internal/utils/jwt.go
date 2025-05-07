package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yourorg/recipe-app/internal/models"
)

// Claims represents the JWT claims
type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(user *models.User) (string, error) {
	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret" // Only for development
	}

	// Set expiration time (24 hours from now)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "recipe-app",
			Subject:   user.Email,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the signed token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_jwt_secret" // Only for development
	}

	// Parse the JWT string and store the result in `claims`
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}