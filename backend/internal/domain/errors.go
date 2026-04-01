package domain

import "errors"

var (
	// ErrEmailRequired is returned when an empty email is provided.
	ErrEmailRequired = errors.New("email is required")

	// ErrPasswordRequired is returned when an empty password is provided.
	ErrPasswordRequired = errors.New("password is required")
)

// UserAlreadyExistsError indicates a registration attempt with a duplicate email.
type UserAlreadyExistsError struct {
	Email string
}

func (e *UserAlreadyExistsError) Error() string {
	return "user with email " + e.Email + " already exists"
}
