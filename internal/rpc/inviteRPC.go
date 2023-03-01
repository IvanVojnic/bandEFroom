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
	SendInvite(ctx context.Context, userCreatorID uuid.UUID, users []models.User, place string, date time.Time) error
	AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, status int) error
	DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, status int) error
}

type InviteServer struct {
	pr.UnimplementedRoomServer
	inviteServ Invite
}

func NewInviteServer(inviteServ Invite) *InviteServer {
	return &InviteServer{inviteServ: inviteServ}
}

func (s *InviteServer) SendInvite(ctx context.Context, req *pr.SendInviteRequest) (*pr.SendInviteResponse, error) {
	userCreatorID, errParse := uuid.Parse(req.GetUserCreatorID())
	if errParse != nil {
		logrus.WithFields(logrus.Fields{
			"Error parse user creator ID": errParse,
			"userCreatorID":               userCreatorID,
		}).Errorf("error parsing ID, %s", errParse)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errParse)
	}
	var users []models.User
	for _, userGRPC := range req.GetFriends() {
		userID, errParseID := uuid.Parse(userGRPC.ID)
		if errParseID != nil {
			logrus.WithFields(logrus.Fields{
				"Error parse user ID": errParseID,
				"userID":              userID,
				"user":                userGRPC,
			}).Errorf("error parsing ID, %s", errParseID)
			return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing ID, %s", errParseID)
		}
		user := models.User{ID: userID, Name: userGRPC.Name, Email: userGRPC.Email}
		users = append(users, user)
	}
	date, errDateParse := time.Parse("2006-01-02 15:04:05", req.GetDate())
	if errDateParse != nil {
		logrus.WithFields(logrus.Fields{
			"Error parse date": errDateParse,
			"date":             date,
		}).Errorf("error parsing date, %s", errDateParse)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while parsing date, %s", errDateParse)
	}
	errSend := s.inviteServ.SendInvite(ctx, userCreatorID, users, req.GetPlace(), date)
	if errSend != nil {
		logrus.WithFields(logrus.Fields{
			"Error sending invite": errSend,
		}).Errorf("error sending invite, %s", errSend)
		return &pr.SendInviteResponse{}, fmt.Errorf("error while seding invite, %s", errSend)
	}
	return &pr.SendInviteResponse{}, nil
}
