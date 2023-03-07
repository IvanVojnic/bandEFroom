// Package service define room services methods
package service

import (
	"context"

	"github.com/IvanVojnic/bandEFroom/models"

	"github.com/google/uuid"
)

// Room interface define implemented methods
type Room interface {
	GetRooms(ctx context.Context, user uuid.UUID) (*[]models.Room, error)
	GetRoomUsers(ctx context.Context, roomID uuid.UUID) (*[]uuid.UUID, error)
	GetUsers(ctx context.Context, usersID *[]uuid.UUID) (*[]models.User, error)
}

// RoomServer define service invites
type RoomServer struct {
	roomRepo Room
}

// NewRoomServer used to init service user communicate struct
func NewRoomServer(roomRepo Room) *RoomServer {
	return &RoomServer{roomRepo: roomRepo}
}

// GetRooms used to get rooms by repo
func (s *RoomServer) GetRooms(ctx context.Context, userID uuid.UUID) (*[]models.Room, error) {
	return s.roomRepo.GetRooms(ctx, userID) // get all user rooms
}

// GetRoomsUser used get users of current room
func (s *RoomServer) GetRoomsUser(ctx context.Context, roomID uuid.UUID) (*[]models.User, error) {
	usersID, errUsersID := s.roomRepo.GetRoomUsers(ctx, roomID) // get array of users id who are in current room
	if errUsersID != nil {
		return nil, errUsersID
	}
	users, err := s.roomRepo.GetUsers(ctx, usersID) // get all users by they id from userMS
	if err != nil {
		return nil, err
	}
	return users, nil
}
