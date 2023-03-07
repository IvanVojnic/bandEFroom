// Package rpc define room rpc methods
package rpc

import (
	"context"
	"fmt"

	"github.com/IvanVojnic/bandEFroom/models"
	pr "github.com/IvanVojnic/bandEFroom/proto"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Room is an interface with implemented methods from room service
type Room interface {
	GetRooms(ctx context.Context, user uuid.UUID) (*[]models.Room, error)
	GetRoomsUser(ctx context.Context, roomID uuid.UUID) (*[]models.User, error)
}

// RoomServer used to define room server obj
type RoomServer struct {
	pr.UnimplementedRoomServer
	roomServ Room
}

// NewRoomServer used to init room server obj
func NewRoomServer(roomServ Room) *RoomServer {
	return &RoomServer{roomServ: roomServ}
}

// GetRooms used to get rooms by serv
func (s *RoomServer) GetRooms(ctx context.Context, req *pr.GetRoomsRequest) (*pr.GetRoomsResponse, error) {
	userID, errUserParse := uuid.Parse(req.GetUserID())
	if errUserParse != nil {
		logrus.WithFields(logrus.Fields{
			"userID": userID,
		}).Errorf("error parsing ID (get rooms), %s", errUserParse)
		return &pr.GetRoomsResponse{}, fmt.Errorf("error while parsing ID, %s", errUserParse)
	}
	rooms, errGetRooms := s.roomServ.GetRooms(ctx, userID)
	if errGetRooms != nil {
		logrus.WithFields(logrus.Fields{
			"rooms": rooms,
		}).Errorf("error get rooms, %s", errGetRooms)
		return &pr.GetRoomsResponse{}, fmt.Errorf("error while get rooms, %s", errGetRooms)
	}
	return &pr.GetRoomsResponse{}, nil
}

// GetUsersRoom used to get users from current room by serv
func (s *RoomServer) GetUsersRoom(ctx context.Context, req *pr.GetUsersRoomRequest) (*pr.GetUsersRoomResponse, error) {
	roomID, errRoomParse := uuid.Parse(req.GetRoomID())
	if errRoomParse != nil {
		logrus.WithFields(logrus.Fields{
			"userID": roomID,
		}).Errorf("error parsing ID (GetUsersRoom), %s", errRoomParse)
		return &pr.GetUsersRoomResponse{}, fmt.Errorf("error while parsing ID, %s", errRoomParse)
	}
	users, errGetUsers := s.roomServ.GetRoomsUser(ctx, roomID)
	if errGetUsers != nil {
		logrus.WithFields(logrus.Fields{
			"users": users,
		}).Errorf("error get users from room, %s", errGetUsers)
		return &pr.GetUsersRoomResponse{}, fmt.Errorf("error while getting users from room, %s", errGetUsers)
	}
	usersGRPC := make([]*pr.User, 0)
	for _, user := range *users {
		usersGRPC = append(usersGRPC, &pr.User{ID: user.ID.String(), Name: user.Name, Email: user.Email})
	}
	return &pr.GetUsersRoomResponse{Users: usersGRPC}, nil
}
