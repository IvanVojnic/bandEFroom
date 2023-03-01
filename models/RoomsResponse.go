package models

import (
	"github.com/google/uuid"
	"time"
)

// RoomResponse is a RoomResponse
type RoomResponse struct {
	ID            uuid.UUID `json:"id" db:"id"`
	UserCreatorID uuid.UUID `json:"userCreatorID" db:"idUserCreator"`
	Place         string    `json:"place" db:"place"`
	Date          time.Time `json:"date" db:"date"`
	Users         []User    `json:"users"`
}
