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

// PasswordHasher defines the port for password hashing.
type PasswordHasher interface {
	Hash(password string) (string, error)
}

// UserService provides user-related business operations.
type UserService struct {
	repo   UserRepository
	hasher PasswordHasher
}

// NewUserService creates a new UserService.
func NewUserService(repo UserRepository, hasher PasswordHasher) *UserService {
	return &UserService{repo: repo, hasher: hasher}
}

// Register handles user registration.
func (s *UserService) Register(ctx context.Context, email, password string) (*domain.User, error) {
	// 1. Validate raw inputs before any transformation
	if email == "" {
		return nil, domain.ErrEmailRequired
	}
	if password == "" {
		return nil, domain.ErrPasswordRequired
	}

	// 2. Check if user already exists
	existing, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existing != nil {
		return nil, &domain.UserAlreadyExistsError{Email: email}
	}

	// 3. Hash password
	hashedPassword, err := s.hasher.Hash(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 4. Create new user (Domain logic)
	user, err := domain.NewUser(email, hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// 5. Save to database
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}
