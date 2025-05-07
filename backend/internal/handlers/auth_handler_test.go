package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheAlonso95/recipe-app/internal/models"
	"github.com/TheAlonso95/recipe-app/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepository is a mock implementation of the UserRepository for testing
type MockUserRepository struct {
	users map[string]*models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*models.User),
	}
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	if _, exists := m.users[user.Email]; exists {
		return repository.ErrDuplicateEmail
	}

	// Auto-increment ID (simple simulation)
	user.ID = len(m.users) + 1
	m.users[user.Email] = user
	return nil
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	user, exists := m.users[email]
	if !exists {
		return nil, repository.ErrUserNotFound
	}
	return user, nil
}

// Mocked AuthHandler for testing
type MockAuthHandler struct {
	AuthHandler
}

func NewMockAuthHandler() *MockAuthHandler {
	mockRepo := NewMockUserRepository()
	return &MockAuthHandler{
		AuthHandler: AuthHandler{
			UserRepo: mockRepo,
		},
	}
}

func TestRegister(t *testing.T) {
	handler := NewMockAuthHandler()

	// Test valid registration
	reqBody := models.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.Register(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response models.AuthResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Token == "" {
		t.Error("Expected token to be returned")
	}

	if response.User.Email != reqBody.Email {
		t.Errorf("Expected email %s, got %s", reqBody.Email, response.User.Email)
	}

	// Test duplicate email
	req = httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewReader(body))
	w = httptest.NewRecorder()

	handler.Register(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("Expected status code %d, got %d", http.StatusConflict, w.Code)
	}
}

func TestLogin(t *testing.T) {
	handler := NewMockAuthHandler()

	// First register a test user
	email := "login-test@example.com"
	password := "password123"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	handler.UserRepo.CreateUser(user)

	// Test valid login
	reqBody := models.LoginRequest{
		Email:    email,
		Password: password,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.Login(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response models.AuthResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Token == "" {
		t.Error("Expected token to be returned")
	}

	// Test invalid password
	reqBody.Password = "wrongpassword"
	body, _ = json.Marshal(reqBody)

	req = httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewReader(body))
	w = httptest.NewRecorder()

	handler.Login(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
	}
}
