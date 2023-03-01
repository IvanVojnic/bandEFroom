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
func (r *RoomPostgres) GetRooms(ctx context.Context, userID uuid.UUID) ([]models.Room, error) {
	var rooms []models.Room
	rowsUserRooms, errUserRooms := r.db.Query(ctx,
		`SELECT ROOMS.id, ROOMS.room_idUserCreator, ROOMS.room_place, ROOMS.room_date, USERS.email, STATUSES.status FROM ROOMS
				INNER JOIN INVITES on INVITES.room_id = ROOMS.id AND INVITES.user_id = $1
				INNER JOIN USERS on USERS.id = ROOMS.room_idUserCreator
				INNER JOIN statuses on STATUSES.id = INVITES.status_id`, userID)
	if errUserRooms != nil {
		return rooms, fmt.Errorf("error while getting invites")
	}
	defer rowsUserRooms.Close()
	for rowsUserRooms.Next() {
		var room models.Room
		errScan := rowsUserRooms.Scan(&room.RoomID, &room.RoomUserID, &room.RoomPlace, &room.RoomDate)
		if errScan != nil {
			return rooms, fmt.Errorf("get all friends requests scan rows error %w", errScan)
		}
		rooms = append(rooms, room)
	}
	for i := 0; i < len(rooms); i++ {
		var users []models.User
		rowsUsers, errUsers := r.db.Query(ctx,
			`SELECT users.id, users.email, invites.status_id FROM users
				INNER JOIN invites on invites.user_id = users.id
				WHERE invites.room_id=$1`, rooms[i].RoomID)
		if errUsers != nil {
			return rooms, fmt.Errorf("error while getting invited users")
		}
		defer rowsUsers.Close()
		for rowsUsers.Next() {
			var user models.User
			errScan := rowsUserRooms.Scan(&user.UserID, &user.UserEmail)
			if errScan != nil {
				return rooms, fmt.Errorf("get all friends requests scan rows error %w", errScan)
			}
			users = append(users, user)
		}
		rooms[i].Users = users
	}
	return rooms, nil
}
