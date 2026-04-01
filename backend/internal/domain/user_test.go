package domain

import (
	"errors"
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		expectError error
	}{
		{
			name:     "Valid user",
			email:    "test@example.com",
			password: "securepassword",
		},
		{
			name:        "Empty email",
			email:       "",
			password:    "securepassword",
			expectError: ErrEmailRequired,
		},
		{
			name:        "Empty password",
			email:       "test@example.com",
			password:    "",
			expectError: ErrPasswordRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.email, tt.password)

			if tt.expectError != nil {
				if !errors.Is(err, tt.expectError) {
					t.Errorf("Expected error %v, got: %v", tt.expectError, err)
				}
				if user != nil {
					t.Errorf("Expected nil user, got %v", user)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
				if user == nil {
					t.Fatalf("Expected valid user, got nil")
				}
				if user.Email != tt.email {
					t.Errorf("Expected email %s, got %s", tt.email, user.Email)
				}
				if user.Password != tt.password {
					t.Errorf("Expected password %s, got %s", tt.password, user.Password)
				}
				if user.ID.String() == "" {
					t.Errorf("Expected valid UUID for user ID, got empty string")
				}
			}
		})
	}
}
