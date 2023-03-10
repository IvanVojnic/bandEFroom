// Package service define invite services methods
package service

import (
	"context"
	"fmt"
	"github.com/IvanVojnic/bandEFroom/models"
	"time"

	"github.com/google/uuid"
)

// Invite interface consists of implemented methods to invite users
type Invite interface {
	SendInvite(ctx context.Context, users []*uuid.UUID, roomID, creatorID uuid.UUID) error
	AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error
	DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error
	CreateRoom(ctx context.Context, userCreatorID uuid.UUID, place string, date time.Time) (uuid.UUID, error)
}

type NotificationMS interface {
	StorageInvite(ctx context.Context, userCreator models.User, usersInvited []*models.User, roomID uuid.UUID, date time.Time, place string) error
}

// InviteServer define service invites
type InviteServer struct {
	inviteRepo     Invite
	userRepo       User
	notificationMS NotificationMS
}

// NewInviteServer used to init service user communicate struct
func NewInviteServer(inviteRepo Invite, userRepo User, notificationMS NotificationMS) *InviteServer {
	return &InviteServer{inviteRepo: inviteRepo, userRepo: userRepo, notificationMS: notificationMS}
}

// SendInvite used to send invite friends to the room by repo
func (s *InviteServer) SendInvite(ctx context.Context, userCreatorID uuid.UUID, usersID []*uuid.UUID, place string, date time.Time) error {
	roomID, err := s.inviteRepo.CreateRoom(ctx, userCreatorID, place, date)
	if err != nil {
		return fmt.Errorf("error while creating room, %s", err)
	}
	err = s.inviteRepo.SendInvite(ctx, usersID, roomID, userCreatorID)
	if err != nil {
		fmt.Errorf("error while creating invites, %s", err)
	}
	userCreator, err := s.userRepo.GetUser(ctx, userCreatorID)
	if err != nil {
		fmt.Errorf("error while getting user creator, %s", err)
	}
	usersInvited, err := s.userRepo.GetUsers(ctx, usersID)
	if err != nil {
		fmt.Errorf("error while getting users, %s", err)
	}
	return s.notificationMS.StorageInvite(ctx, userCreator, usersInvited, roomID, date, place)
}

// AcceptInvite used to accept invite to the room by repo
func (s *InviteServer) AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error {
	return s.inviteRepo.AcceptInvite(ctx, userID, roomID)
}

// DeclineInvite used to decline invite to the room by repo
func (s *InviteServer) DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error {
	return s.inviteRepo.DeclineInvite(ctx, userID, roomID)
}
