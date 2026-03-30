// Package api contains the HTTP handlers and DTOs.
package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/myorg/myapp/backend/internal/api/dtos"
	"github.com/myorg/myapp/backend/internal/application"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userService *application.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register handles the POST /register endpoint.
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dtos.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dtos.UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}
