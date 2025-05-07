package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/yourorg/recipe-app/internal/models"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyTaken = errors.New("email already taken")
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

// PostgresUserRepository implements UserRepository for PostgreSQL
type PostgresUserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{
		DB: db,
	}
}

// CreateUser creates a new user in the database
func (r *PostgresUserRepository) CreateUser(user *models.User) error {
	// Check if user with same email already exists
	existingUser, err := r.GetUserByEmail(user.Email)
	if err != nil && err != ErrUserNotFound {
		return err
	}
	if existingUser != nil {
		return ErrEmailAlreadyTaken
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Insert user into database
	query := `
		INSERT INTO users (email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err = r.DB.QueryRow(
		query,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)

	return err
}

// GetUserByEmail retrieves a user by their email address
func (r *PostgresUserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password, created_at, updated_at
		FROM users
		WHERE email = $1`

	user := &models.User{}
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}