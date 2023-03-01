package service

import (
	"context"
	"fmt"
	"github.com/IvanVojnic/bandEFroom/models"
	"github.com/google/uuid"
)

type Room interface {
	GetRooms(ctx context.Context, user uuid.UUID) ([]models.Room, error)
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
func (s *RoomServer) GetRooms(ctx context.Context, userID uuid.UUID) ([]models.RoomResponse, error) {
	rooms, err := s.roomRepo.GetRooms(ctx, userID)
	fmt.Printf("room %s", rooms)
	var testRooms []models.RoomResponse
	if err != nil {
		return testRooms, fmt.Errorf("error while creating room, %s", err)
	}
	return testRooms, nil
}
