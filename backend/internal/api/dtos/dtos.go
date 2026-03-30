// Package dtos contains the Data Transfer Objects for the API.
package dtos

import "time"

// RegisterRequest is the payload for user registration.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse is the public representation of a user.
type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
