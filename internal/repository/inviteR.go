package repository

import (
	"context"
	"fmt"
	"github.com/IvanVojnic/bandEFroom/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// InvitePostgres is a wrapper to db object
type InvitePostgres struct {
	db *pgxpool.Pool
}

// Status used to define types of statuses
type Status int

// Decline define negative status
// Accept define positive status
// NoAnswer define neutral status
const (
	Decline Status = iota
	Accept
	NoAnswer
)

// NewInvitePostgres used to init InvitePostgres
func NewInvitePostgres(db *pgxpool.Pool) *InvitePostgres {
	return &InvitePostgres{db: db}
}

// SendInvite used to send request to be a friends
func (r *RoomPostgres) SendInvite(ctx context.Context, users []models.User, roomID uuid.UUID, creatorID uuid.UUID) error {
	for i := 0; i < len(users); i++ {
		inviteID := uuid.New()
		_, errInvite := r.db.Exec(ctx, "insert into invites (id, user_id, room_id, status) values($1, $2, $3, $4)",
			inviteID, users[i].ID, roomID, NoAnswer)
		if errInvite != nil {
			return fmt.Errorf("error while invite creating: %s", errInvite)
		}
	}
	inviteID := uuid.New()
	_, errInvite := r.db.Exec(ctx, "insert into invites (id, user_id, room_id, status) values($1, $2, $3, $4)",
		inviteID, creatorID, roomID, Accept)
	if errInvite != nil {
		return fmt.Errorf("error while invite creating: %s", errInvite)
	}
	return nil
}

// AcceptInvite used to accept invite to the room
func (r *RoomPostgres) AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, statusID int) error {
	_, err := r.db.Exec(ctx,
		`UPDATE invites 
			SET status_id=$1 
			WHERE invites.user_id=$2 AND invites.room_id=$3`,
		1, userID, roomID)
	if err != nil {
		return fmt.Errorf("update friends error %w", err)
	}
	return nil
}
