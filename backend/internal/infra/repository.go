// Package infra contains the technical implementations of the application ports.
package infra

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/myorg/myapp/backend/internal/application"
	"github.com/myorg/myapp/backend/internal/domain"
	"github.com/myorg/myapp/backend/internal/infra/database"
)

// SQLiteUserRepository implements application.UserRepository using SQLite.
type SQLiteUserRepository struct {
	queries *database.Queries
	db      *sql.DB
}

// NewSQLiteUserRepository creates a new SQLiteUserRepository.
func NewSQLiteUserRepository(db *sql.DB) application.UserRepository {
	return &SQLiteUserRepository{
		queries: database.New(db),
		db:      db,
	}
}

// Create inserts a new user into the database.
func (r *SQLiteUserRepository) Create(ctx context.Context, user *domain.User) error {
	params := database.CreateUserParams{
		ID:        user.ID.String(),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if err := r.queries.CreateUser(ctx, params); err != nil {
		return fmt.Errorf("failed to create user in DB: %w", err)
	}
	return nil
}

// GetByEmail retrieves a user by email.
func (r *SQLiteUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email from DB: %w", err)
	}

	return toDomain(u)
}

// GetByID retrieves a user by ID.
func (r *SQLiteUserRepository) GetByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	u, err := r.queries.GetUserByID(ctx, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id from DB: %w", err)
	}

	return toDomain(u)
}

func toDomain(u database.User) (*domain.User, error) {
	userID, err := domain.ParseUserID(u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID: %w", err)
	}

	return &domain.User{
		ID:        userID,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
