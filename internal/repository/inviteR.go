// Package repository define invite repo methods
package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/IvanVojnic/bandEFroom/models"

	prNotif "github.com/IvanVojnic/bandEFnotif/proto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// InvitePostgres is a wrapper to db object
type InvitePostgres struct {
	db     *pgxpool.Pool
	client prNotif.InviteRoomClient
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
func (r *InvitePostgres) SendInvite(ctx context.Context, users []*models.User, roomID, creatorID uuid.UUID) error {
	batch := &pgx.Batch{}
	var inviteID uuid.UUID
	for _, user := range users {
		inviteID = uuid.New()
		batch.Queue("INSERT INTO invites (id, user_id, room_id, status) VALUES($1, $2, $3, $4)", inviteID, user.ID, roomID, NoAnswer)
	}
	inviteID = uuid.New()
	batch.Queue("INSERT INTO invites (id, user_id, room_id, status) VALUES($1, $2, $3, $4)", inviteID, creatorID, roomID, Accept)
	res := r.db.SendBatch(ctx, batch)
	err := res.Close()
	if err != nil {
		return fmt.Errorf("error while inserting into invites: %s", err)
	}
	return nil
}

// AcceptInvite used to accept invite to the room
func (r *InvitePostgres) AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error {
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
func (r *InvitePostgres) DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error {
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
func (r *InvitePostgres) CreateRoom(ctx context.Context, userID uuid.UUID, place string, date time.Time) (uuid.UUID, error) {
	roomID := uuid.New()
	_, errRoom := r.db.Exec(ctx, "INSERT INTO rooms (id, idUserCreator, place, date) VALUES($1, $2, $3, $4)",
		roomID, userID, place, date)
	if errRoom != nil {
		return uuid.UUID{}, fmt.Errorf("error while room creating: %s", errRoom)
	}
	return roomID, nil
}

func (r *InvitePostgres) StorageInvite(ctx context.Context, userCreatorID uuid.UUID, roomID uuid.UUID, date time.Time, place string) error {
	_, errGRPC := r.client.StorageInvite(ctx, &prNotif.StorageInviteRequest{UserCreatorID: userCreatorID.String(), RoomID: roomID.String(), Place: place, Date: date.String()})
	if errGRPC != nil {
		return fmt.Errorf("error while storage notiffications of invite, %s", errGRPC)
	}
	return nil
}
