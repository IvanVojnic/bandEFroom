package repository

import (
	"context"
	"fmt"
	"github.com/IvanVojnic/bandEFroom/models"
	"github.com/jackc/pgx/v5"
	"time"

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
func (r *RoomPostgres) SendInvite(ctx context.Context, users *[]models.User, roomID uuid.UUID, creatorID uuid.UUID) error {
	batch := &pgx.Batch{}
	for _, user := range *users {
		inviteID := uuid.New()
		batch.Queue("INSERT INTO invites (id, user_id, room_id, status) VALUES($1, $2, $3, $4)", inviteID, user.ID, roomID, NoAnswer)
	}
	res := r.db.SendBatch(ctx, batch)
	defer res.Close()
	inviteID := uuid.New()
	_, errInvite := r.db.Exec(ctx, "insert into invites (id, user_id, room_id, status) values($1, $2, $3, $4)",
		inviteID, creatorID, roomID, Accept)
	if errInvite != nil {
		return fmt.Errorf("error while invite creating: %s", errInvite)
	}
	return nil
}

// AcceptInvite used to accept invite to the room
func (r *RoomPostgres) AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {
	_, err := r.db.Exec(ctx,
		`UPDATE invites 
			SET status=$1 
			WHERE invites.user_id=$2 AND invites.room_id=$3`,
		Accept, userID, roomID)
	if err != nil {
		return fmt.Errorf("accpet invite error %w", err)
	}
	return nil
}

// DeclineInvite used to accept invite to the room
func (r *RoomPostgres) DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {
	_, err := r.db.Exec(ctx,
		`UPDATE invites 
			SET status=$1 
			WHERE invites.user_id=$2 AND invites.room_id=$3`,
		Decline, userID, roomID)
	if err != nil {
		return fmt.Errorf("accpet invite error %w", err)
	}
	return nil
}

// CreateRoom used to create room
func (r *RoomPostgres) CreateRoom(ctx context.Context, userID uuid.UUID, place string, date time.Time) (uuid.UUID, error) {
	roomID := uuid.New()
	_, errRoom := r.db.Exec(ctx, "insert into rooms (id, idUserCreator, place, date) values($1, $2, $3, $4)",
		roomID, userID, place, date)
	if errRoom != nil {
		return uuid.UUID{}, fmt.Errorf("error while room creating: %s", errRoom)
	}
	return roomID, nil
}
