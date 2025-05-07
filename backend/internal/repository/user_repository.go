package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/TheAlonso95/recipe-app/internal/models"
)

var (
	// ErrDuplicateEmail is returned when trying to create a user with an email that already exists
	ErrDuplicateEmail = errors.New("email already exists")
	// ErrUserNotFound is returned when a user is not found in the database
	ErrUserNotFound = errors.New("user not found")
)

// UserRepository handles database operations related to users
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser adds a new user to the database
func (r *UserRepository) CreateUser(user *models.User) error {
	// Check if user with this email already exists
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", user.Email).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking email existence: %w", err)
	}

	if exists {
		return ErrDuplicateEmail
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Insert the new user
	query := `
    INSERT INTO users (email, password, created_at, updated_at)
    VALUES ($1, $2, $3, $4)
    RETURNING id`

	err = r.DB.QueryRow(query, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

// GetUserByEmail retrieves a user by email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
    SELECT id, email, password, created_at, updated_at
    FROM users
    WHERE email = $1`

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("error getting user by email: %w", err)
	}

	return user, nil
}
