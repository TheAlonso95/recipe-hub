package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yourorg/recipe-app/internal/models"
	"github.com/yourorg/recipe-app/internal/repositories"
	"github.com/yourorg/recipe-app/internal/utils"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	UserRepo repositories.UserRepository
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(userRepo repositories.UserRepository) *AuthHandler {
	return &AuthHandler{
		UserRepo: userRepo,
	}
}

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the response for authentication operations
type AuthResponse struct {
	Token string `json:"token"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		sendError(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		sendError(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Create user
	user := &models.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Save user to database
	err = h.UserRepo.CreateUser(user)
	if err != nil {
		if errors.Is(err, repositories.ErrEmailAlreadyTaken) {
			sendError(w, "Email already registered", http.StatusConflict)
		} else {
			sendError(w, "Error creating user", http.StatusInternalServerError)
		}
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		sendError(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return token
	sendJSON(w, AuthResponse{Token: token}, http.StatusCreated)
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		sendError(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Get user from database
	user, err := h.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			sendError(w, "Invalid email or password", http.StatusUnauthorized)
		} else {
			sendError(w, "Error retrieving user", http.StatusInternalServerError)
		}
		return
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		sendError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		sendError(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return token
	sendJSON(w, AuthResponse{Token: token}, http.StatusOK)
}

// Helper function to send JSON response
func sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Helper function to send error response
func sendError(w http.ResponseWriter, message string, statusCode int) {
	sendJSON(w, ErrorResponse{Message: message}, statusCode)
}