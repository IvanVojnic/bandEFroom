// Package rpc define invite rpc methods
package rpc

import (
	"context"
	"fmt"
	"time"

	pr "github.com/IvanVojnic/bandEFroom/proto"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Invite is an interface with implemented methods from Invite service
type Invite interface {
	SendInvite(ctx context.Context, userCreatorID uuid.UUID, users []*uuid.UUID, place string, date time.Time) error
	AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error
	DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error
}

// InviteServer used to define invite server obj
type InviteServer struct {
	pr.UnimplementedInviteServer
	inviteServ Invite
}

// timeLayout define layout for date parsing
const timeLayout = "2006-01-02 15:04:05"

// NewInviteServer used to init invite serv obj
func NewInviteServer(inviteServ Invite) *InviteServer {
	return &InviteServer{inviteServ: inviteServ}
}

// SendInvite used to send invite by serv
func (s *InviteServer) SendInvite(ctx context.Context, req *pr.SendInviteRequest) (*pr.SendInviteResponse, error) {
	userCreatorID, errParse := uuid.Parse(req.GetUserCreatorID())
	if errParse != nil {
		logrus.WithFields(logrus.Fields{
			"userCreatorID": userCreatorID,
		}).Errorf("error parsing ID (send invite), %s", errParse)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errParse)
	}
	usersID := make([]*uuid.UUID, len(req.GetUsersID()))
	for _, userGRPC := range req.GetUsersID() {
		userID, errParseID := uuid.Parse(userGRPC)
		if errParseID != nil {
			logrus.WithFields(logrus.Fields{
				"userID": userID,
				"user":   userGRPC,
			}).Errorf("error parsing ID (send invite), %s", errParseID)
			return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errParseID)
		}
		usersID = append(usersID, &userID)
	}
	date, errDateParse := time.Parse(timeLayout, req.GetDate())
	if errDateParse != nil {
		logrus.WithFields(logrus.Fields{
			"date": date,
		}).Errorf("error parsing date (send invite), %s", errDateParse)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing date, %s", errDateParse)
	}
	errSend := s.inviteServ.SendInvite(ctx, userCreatorID, usersID, req.GetPlace(), date)
	if errSend != nil {
		logrus.Errorf("error sending invite, %s", errSend)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while seding invite, %s", errSend)
	}
	return &pr.SendInviteResponse{}, nil
}

// AcceptInvite used to accept invite by serv
func (s *InviteServer) AcceptInvite(ctx context.Context, req *pr.AcceptInviteRequest) (*pr.AcceptInviteResponse, error) { // nolint:dupl, gocritic
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

// DeclineInvite used to decline invite by serv
func (s *InviteServer) DeclineInvite(ctx context.Context, req *pr.DeclineInviteRequest) (*pr.DeclineInviteResponse, error) { // nolint:dupl, gocritic
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
