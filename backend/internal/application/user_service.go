// Package application contains the use cases and repository interfaces.
package application

import (
	"context"
	"fmt"

	"github.com/myorg/myapp/backend/internal/domain"
)

// UserRepository defines the persistence port for User entities.
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByID(ctx context.Context, id domain.UserID) (*domain.User, error)
}

// UserService provides user-related business operations.
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Register handles user registration.
func (s *UserService) Register(ctx context.Context, email, password string) (*domain.User, error) {
	// 1. Check if user already exists
	existing, err := s.repo.GetByEmail(ctx, email)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	// 2. Create new user (Domain logic)
	user, err := domain.NewUser(email, password) // Password should be hashed in production
	if err != nil {
		return nil, fmt.Errorf("failed to create user domain object: %w", err)
	}

	// 3. Save to database
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}
