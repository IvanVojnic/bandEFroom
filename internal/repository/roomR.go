// Package repository define room repo methods
package repository

import (
	"context"
	"fmt"

	"github.com/IvanVojnic/bandEFroom/models"
	pr "github.com/IvanVojnic/bandEFuser/proto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RoomPostgres is a wrapper to db object
type RoomPostgres struct {
	db     *pgxpool.Pool
	client pr.UserCommClient
}

// NewRoomPostgres used to init RoomPostgres
func NewRoomPostgres(db *pgxpool.Pool, client pr.UserCommClient) *RoomPostgres {
	return &RoomPostgres{db: db, client: client}
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

// GetUsers used to get full info of users from userMS
func (r *RoomPostgres) GetUsers(ctx context.Context, usersID []*uuid.UUID) ([]*models.User, error) {
	users := make([]*models.User, 0)
	usersIDStr := make([]string, 0)
	for _, ID := range usersID {
		usersIDStr = append(usersIDStr, ID.String())
	}

	res, errGRPC := r.client.GetUsers(ctx, &pr.GetUsersRequest{UsersID: usersIDStr})
	if errGRPC != nil {
		return users, fmt.Errorf("error while sign up, %s", errGRPC)
	}

	for _, user := range res.Users {
		userID, errUserID := uuid.Parse(user.ID)
		if errUserID != nil {
			return users, fmt.Errorf("error while parsing room ID, %s", errUserID)
		}
		users = append(users, &models.User{ID: userID, Name: user.Name, Email: user.Email})
	}
	return users, nil
}
