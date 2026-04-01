// Package api contains the HTTP handlers and DTOs.
package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/myorg/myapp/backend/internal/api/dtos"
	"github.com/myorg/myapp/backend/internal/application"
	"github.com/myorg/myapp/backend/internal/domain"
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
		var alreadyExists *domain.UserAlreadyExistsError
		switch {
		case errors.As(err, &alreadyExists):
			http.Error(w, alreadyExists.Error(), http.StatusConflict)
		case errors.Is(err, domain.ErrEmailRequired) || errors.Is(err, domain.ErrPasswordRequired):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			slog.Error("registration failed", "error", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
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
		slog.Error("failed to encode response", "error", err)
	}
}
