package rpc

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	pr "github.com/IvanVojnic/bandEFroom/proto"

	"github.com/IvanVojnic/bandEFroom/models"
	"github.com/google/uuid"
)

type Invite interface {
	SendInvite(ctx context.Context, userCreatorID uuid.UUID, users *[]models.User, place string, date time.Time) error
	AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error
	DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error
}

type InviteServer struct {
	pr.UnimplementedRoomServer
	inviteServ Invite
}

const timeLayout = "2006-01-02 15:04:05"

func NewInviteServer(inviteServ Invite) *InviteServer {
	return &InviteServer{inviteServ: inviteServ}
}

func (s *InviteServer) SendInvite(ctx context.Context, req *pr.SendInviteRequest) (*pr.SendInviteResponse, error) {
	userCreatorID, errParse := uuid.Parse(req.GetUserCreatorID())
	if errParse != nil {
		logrus.WithFields(logrus.Fields{
			"userCreatorID": userCreatorID,
		}).Errorf("error parsing ID (send invite), %s", errParse)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errParse)
	}
	users := make([]models.User, len(req.GetUsersID()))
	for _, userGRPC := range req.GetUsersID() {
		userID, errParseID := uuid.Parse(userGRPC)
		if errParseID != nil {
			logrus.WithFields(logrus.Fields{
				"userID": userID,
				"user":   userGRPC,
			}).Errorf("error parsing ID (send invite), %s", errParseID)
			return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errParseID)
		}
		user := models.User{ID: userID}
		users = append(users, user)
	}
	date, errDateParse := time.Parse(timeLayout, req.GetDate())
	if errDateParse != nil {
		logrus.WithFields(logrus.Fields{
			"date": date,
		}).Errorf("error parsing date (send invite), %s", errDateParse)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing date, %s", errDateParse)
	}
	errSend := s.inviteServ.SendInvite(ctx, userCreatorID, &users, req.GetPlace(), date)
	if errSend != nil {
		logrus.Errorf("error sending invite, %s", errSend)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while seding invite, %s", errSend)
	}
	return &pr.SendInviteResponse{}, nil
}

func (s *InviteServer) AcceptInvite(ctx context.Context, req *pr.AcceptInviteRequest) (*pr.AcceptInviteResponse, error) {
	userID, errUserParse := uuid.Parse(req.GetUserID())
	if errUserParse != nil {
		logrus.WithFields(logrus.Fields{
			"userID": userID,
		}).Errorf("error parsing ID (accept invite), %s", errUserParse)
		return &pr.AcceptInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errUserParse)
	}
	roomID, errRoomParse := uuid.Parse(req.GetRoomID())
	if errRoomParse != nil {
		logrus.WithFields(logrus.Fields{
			"roomID": roomID,
		}).Errorf("error parsing ID (accept invite), %s", errRoomParse)
		return &pr.AcceptInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errRoomParse)
	}
	errAccept := s.inviteServ.AcceptInvite(ctx, userID, roomID)
	if errAccept != nil {
		logrus.Errorf("error accept invite, %s", errAccept)
		return &pr.AcceptInviteResponse{}, fmt.Errorf("error while accepting invite, %s", errAccept)
	}
	return &pr.AcceptInviteResponse{}, nil
}

func (s *InviteServer) DeclineInvite(ctx context.Context, req *pr.DeclineInviteRequest) (*pr.DeclineInviteResponse, error) {
	userID, errUserParse := uuid.Parse(req.GetUserID())
	if errUserParse != nil {
		logrus.WithFields(logrus.Fields{
			"userID": userID,
		}).Errorf("error parsing ID (decline invite), %s", errUserParse)
		return &pr.DeclineInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errUserParse)
	}
	roomID, errRoomParse := uuid.Parse(req.GetRoomID())
	if errRoomParse != nil {
		logrus.WithFields(logrus.Fields{
			"roomID": roomID,
		}).Errorf("error parsing ID (decline invite), %s", errRoomParse)
		return &pr.DeclineInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errRoomParse)
	}
	errDecline := s.inviteServ.AcceptInvite(ctx, userID, roomID)
	if errDecline != nil {
		logrus.Errorf("error decline invite, %s", errDecline)
		return &pr.DeclineInviteResponse{}, fmt.Errorf("error while decling invite, %s", errDecline)
	}
	return &pr.DeclineInviteResponse{}, nil
}
