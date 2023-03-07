// Package models define RoomResponse model
package models

import (
	"time"

	"github.com/google/uuid"
)

// RoomResponse is a RoomResponse
type RoomResponse struct {
	ID            uuid.UUID `json:"id" db:"id"`
	UserCreatorID uuid.UUID `json:"userCreatorID" db:"idUserCreator"`
	Place         string    `json:"place" db:"place"`
	Date          time.Time `json:"date" db:"date"`
	Users         []User    `json:"users"`
}
