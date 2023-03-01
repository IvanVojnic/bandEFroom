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
}

type RoomServer struct {
	pr.UnimplementedRoomServer
	roomServ Room
}

func NewRoomServer(roomServ Room) *RoomServer {
	return &RoomServer{roomServ: roomServ}
}

func (s *RoomServer) DeclineInvite(ctx context.Context, req *pr.DeclineInviteRequest) (*pr.DeclineInviteResponse, error) {
	userID, errUserParse := uuid.Parse(req.GetUserID())
	if errUserParse != nil {
		logrus.WithFields(logrus.Fields{
			"Error parse user ID": errUserParse,
			"userID":              userID,
		}).Errorf("error parsing ID, %s", errUserParse)
		return &pr.DeclineInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errUserParse)
	}
	roomID, errRoomParse := uuid.Parse(req.GetRoomID())
	if errRoomParse != nil {
		logrus.WithFields(logrus.Fields{
			"Error parse room ID": errRoomParse,
			"roomID":              roomID,
		}).Errorf("error parsing ID, %s", errRoomParse)
		return &pr.DeclineInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errRoomParse)
	}
	errDecline := s.inviteServ.AcceptInvite(ctx, userID, roomID, int(req.GetStatusID()))
	if errDecline != nil {
		logrus.WithFields(logrus.Fields{
			"Error decline invite": errDecline,
		}).Errorf("error decline invite, %s", errDecline)
		return &pr.DeclineInviteResponse{}, fmt.Errorf("error while decling invite, %s", errDecline)
	}
	return &pr.DeclineInviteResponse{}, nil
}
