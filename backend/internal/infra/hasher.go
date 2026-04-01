package infra

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/myorg/myapp/backend/internal/application"
)

// BcryptPasswordHasher implements application.PasswordHasher using bcrypt.
type BcryptPasswordHasher struct {
	cost int
}

// NewBcryptPasswordHasher creates a new BcryptPasswordHasher with default cost.
func NewBcryptPasswordHasher() application.PasswordHasher {
	return &BcryptPasswordHasher{cost: bcrypt.DefaultCost}
}

// Hash generates a bcrypt hash of the given password.
func (h *BcryptPasswordHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", fmt.Errorf("bcrypt hash failed: %w", err)
	}
	return string(hash), nil
}
