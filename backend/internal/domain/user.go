// Package domain contains the core business entities and logic.
package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// UserID is a custom type for user identifiers.
type UserID uuid.UUID

// String returns the string representation of UserID.
func (id UserID) String() string {
	return uuid.UUID(id).String()
}

// ParseUserID parses a string into a UserID.
func ParseUserID(s string) (UserID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return UserID{}, fmt.Errorf("invalid user id: %w", err)
	}
	return UserID(id), nil
}

// NewUserID generates a new UserID.
func NewUserID() UserID {
	return UserID(uuid.New())
}

// User represents a system user.
type User struct {
	ID        UserID
	Email     string
	Password  string // Hashed
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new User entity.
func NewUser(email, password string) (*User, error) {
	if email == "" {
		return nil, ErrEmailRequired
	}
	if password == "" {
		return nil, ErrPasswordRequired
	}
	now := time.Now()
	return &User{
		ID:        NewUserID(),
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
