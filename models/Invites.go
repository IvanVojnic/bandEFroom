// Package models define Invite model
package models

import (
	"github.com/google/uuid"
)

// Invite is an InviteStruct
type Invite struct {
	ID     uuid.UUID `json:"id" db:"id"`
	UserID uuid.UUID `json:"userID" db:"user_id"`
	RoomID uuid.UUID `json:"roomID" db:"room_id"`
	Status int       `json:"status" db:"status"`
}
