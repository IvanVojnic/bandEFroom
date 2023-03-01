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
	GetRooms(ctx context.Context, user uuid.UUID) ([]models.RoomResponse, error)
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
			"Error decline invite": errGetRooms,
			"rooms":                rooms,
		}).Errorf("error decline invite, %s", errGetRooms)
		return &pr.GetRoomsResponse{}, fmt.Errorf("error while decling invite, %s", errGetRooms)
	}
	return &pr.GetRoomsResponse{}, nil
}
