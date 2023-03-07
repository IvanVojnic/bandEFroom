// Package models define User model
package models

import "github.com/google/uuid"

// User is a User
type User struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Email string    `json:"email" db:"email"`
	Name  string    `json:"name" db:"name"`
}
