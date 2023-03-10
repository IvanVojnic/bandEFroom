// Package repository define room repo methods
package repository

import (
	"context"
	"fmt"

	"github.com/IvanVojnic/bandEFroom/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RoomPostgres is a wrapper to db object
type RoomPostgres struct {
	db *pgxpool.Pool
}

// NewRoomPostgres used to init RoomPostgres
func NewRoomPostgres(db *pgxpool.Pool) *RoomPostgres {
	return &RoomPostgres{db: db}
}

// GetRooms used to get rooms where you had invited
func (r *RoomPostgres) GetRooms(ctx context.Context, userID uuid.UUID) ([]*models.Room, error) {
	var rooms []*models.Room
	rowsUserRooms, errUserRooms := r.db.Query(ctx,
		`SELECT ROOMS.id, ROOMS.idUserCreator, ROOMS.place, ROOMS.date FROM ROOMS
			 INNER JOIN INVITES on INVITES.room_id = ROOMS.id AND INVITES.user_id = $1`, userID)
	if errUserRooms != nil {
		return rooms, fmt.Errorf("error while getting invites, %s", errUserRooms)
	}
	defer rowsUserRooms.Close()

	for rowsUserRooms.Next() {
		var room models.Room
		errScan := rowsUserRooms.Scan(&room.ID, &room.UserCreatorID, &room.Place, &room.Date)
		if errScan != nil {
			return rooms, fmt.Errorf("get all friends requests scan rows error %w", errScan)
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

// GetRoomUsers used to get users from current room
func (r *RoomPostgres) GetRoomUsers(ctx context.Context, roomID uuid.UUID) ([]*uuid.UUID, error) {
	usersID := make([]*uuid.UUID, 0)
	rowsUserInvites, errUserInvites := r.db.Query(ctx,
		`SELECT INVITES.user_id FROM INVITES
			 WHERE INVITES.room_id=$1`, roomID)
	if errUserInvites != nil {
		return usersID, fmt.Errorf("error while getting invites, %s", errUserInvites)
	}
	defer rowsUserInvites.Close()

	for rowsUserInvites.Next() {
		var userID uuid.UUID
		errScan := rowsUserInvites.Scan(&userID)
		if errScan != nil {
			return usersID, fmt.Errorf("get all friends requests scan rows error %w", errScan)
		}
		usersID = append(usersID, &userID)
	}
	return usersID, nil
}
