package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/yourorg/recipe-app/internal/models"
	"github.com/yourorg/recipe-app/internal/repositories"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	users map[string]*models.User
}

// CreateUser is a mock implementation of UserRepository.CreateUser
func (r *MockUserRepository) CreateUser(user *models.User) error {
	// Check if user with same email already exists
	if _, exists := r.users[user.Email]; exists {
		return repositories.ErrEmailAlreadyTaken
	}

	// Set timestamps and mock ID
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.ID = int64(len(r.users) + 1) // Simple ID generation for tests

	// Save user to mock storage
	r.users[user.Email] = user
	return nil
}

// GetUserByEmail is a mock implementation of UserRepository.GetUserByEmail
func (r *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	user, exists := r.users[email]
	if !exists {
		return nil, repositories.ErrUserNotFound
	}
	return user, nil
}

// NewMockUserRepository creates a new mock user repository for testing
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*models.User),
	}
}

func TestRegister(t *testing.T) {
	// Create mock repository and handler
	mockRepo := NewMockUserRepository()
	handler := NewAuthHandler(mockRepo)

	// Create test request
	reqBody := RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.Register(rr, req)

	// Check response status
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check response body contains token
	var response AuthResponse
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.Token == "" {
		t.Error("Expected token in response, got empty string")
	}

	// Test duplicate email registration
	rr = httptest.NewRecorder()
	handler.Register(rr, req)
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("Handler should return conflict on duplicate email: got %v want %v", status, http.StatusConflict)
	}
}

func TestLogin(t *testing.T) {
	// Create mock repository and handler
	mockRepo := NewMockUserRepository()
	handler := NewAuthHandler(mockRepo)

	// Create a user for testing
	user := &models.User{
		Email:    "test@example.com",
		Password: "", // Will be set by CreateUser
	}

	// Hash the password before saving
	password := "password123"
	hashedPassword, _ := utils.HashPassword(password)
	user.Password = hashedPassword
	mockRepo.CreateUser(user)

	// Test successful login
	reqBody := LoginRequest{
		Email:    "test@example.com",
		Password: password,
	}
	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	// Check response status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body contains token
	var response AuthResponse
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response.Token == "" {
		t.Error("Expected token in response, got empty string")
	}

	// Test invalid password
	reqBody.Password = "wrongpassword"
	body, _ = json.Marshal(reqBody)
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()

	handler.Login(rr, req)

	// Check response status
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler should return unauthorized on wrong password: got %v want %v", status, http.StatusUnauthorized)
	}

	// Test invalid email
	reqBody.Email = "nonexistent@example.com"
	body, _ = json.Marshal(reqBody)
	req, _ = http.NewRequest("POST", "/api/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()

	handler.Login(rr, req)

	// Check response status
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler should return unauthorized on wrong email: got %v want %v", status, http.StatusUnauthorized)
	}
}