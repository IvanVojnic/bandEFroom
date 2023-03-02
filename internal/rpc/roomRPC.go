package rpc

import (
	"context"
	"fmt"
	"github.com/IvanVojnic/bandEFroom/models"
	pr "github.com/IvanVojnic/bandEFroom/proto"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type Room interface {
	GetRooms(ctx context.Context, user uuid.UUID) ([]models.Room, error)
	GetUsersRoom(ctx context.Context, roomID uuid.UUID) ([]models.User, error)
}

type RoomServer struct {
	pr.UnimplementedRoomServer
	roomServ Room
}

func NewRoomServer(roomServ Room) *RoomServer {
	return &RoomServer{roomServ: roomServ}
}

func (s *RoomServer) GetRooms(ctx context.Context, req *pr.GetRoomsRequest) (*pr.GetRoomsResponse, error) {
	userID, errUserParse := uuid.Parse(req.GetUserID())
	if errUserParse != nil {
		logrus.WithFields(logrus.Fields{
			"Error parse user ID": errUserParse,
			"userID":              userID,
		}).Errorf("error parsing ID, %s", errUserParse)
		return &pr.GetRoomsResponse{}, fmt.Errorf("error while parsing ID, %s", errUserParse)
	}
	rooms, errGetRooms := s.roomServ.GetRooms(ctx, userID)
	if errGetRooms != nil {
		logrus.WithFields(logrus.Fields{
			"Error get rooms": errGetRooms,
			"rooms":           rooms,
		}).Errorf("error get rooms, %s", errGetRooms)
		return &pr.GetRoomsResponse{}, fmt.Errorf("error while get rooms, %s", errGetRooms)
	}
	return &pr.GetRoomsResponse{}, nil
}

func (s *RoomServer) GetUsersRoom(ctx context.Context, req *pr.GetUsersRoomRequest) (*pr.GetUsersRoomResponse, error) {
	roomID, errRoomParse := uuid.Parse(req.GetRoomID())
	if errRoomParse != nil {
		logrus.WithFields(logrus.Fields{
			"Error parse room ID": errRoomParse,
			"userID":              roomID,
		}).Errorf("error parsing ID, %s", errRoomParse)
		return &pr.GetUsersRoomResponse{}, fmt.Errorf("error while parsing ID, %s", errRoomParse)
	}
	users, errGetUsers := s.roomServ.GetUsersRoom(ctx, roomID)
	if errGetUsers != nil {
		logrus.WithFields(logrus.Fields{
			"Error get users from room": errGetUsers,
			"users":                     users,
		}).Errorf("error get users from room, %s", errGetUsers)
		return &pr.GetUsersRoomResponse{}, fmt.Errorf("error while getting users from room, %s", errGetUsers)
	}
	return &pr.GetUsersRoomResponse{}, nil
}
