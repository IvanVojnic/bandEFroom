// Package service define invite services methods
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/IvanVojnic/bandEFroom/models"

	"github.com/google/uuid"
)

// Invite interface consists of implemented methods to invite users
type Invite interface {
	SendInvite(ctx context.Context, users []*models.User, roomID, creatorID uuid.UUID) error
	AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error
	DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error
	CreateRoom(ctx context.Context, userCreatorID uuid.UUID, place string, date time.Time) (uuid.UUID, error)
	StorageInvite(ctx context.Context, userCreatorID uuid.UUID, roomID uuid.UUID, date time.Time, place string) error
}

// InviteServer define service invites
type InviteServer struct {
	inviteRepo Invite
}

// NewInviteServer used to init service user communicate struct
func NewInviteServer(inviteRepo Invite) *InviteServer {
	return &InviteServer{inviteRepo: inviteRepo}
}

// SendInvite used to send invite friends to the room by repo
func (s *InviteServer) SendInvite(ctx context.Context, userCreatorID uuid.UUID, users []*models.User, place string, date time.Time) error {
	roomID, err := s.inviteRepo.CreateRoom(ctx, userCreatorID, place, date)
	if err != nil {
		return fmt.Errorf("error while creating room, %s", err)
	}
	err = s.inviteRepo.SendInvite(ctx, users, roomID, userCreatorID)
	if err != nil {
		fmt.Errorf("error while creating invites, %s", err)
	}
	return s.inviteRepo.StorageInvite(ctx, userCreatorID, roomID, date, place)
}

// AcceptInvite used to accept invite to the room by repo
func (s *InviteServer) AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error {
	return s.inviteRepo.AcceptInvite(ctx, userID, roomID)
}

// DeclineInvite used to decline invite to the room by repo
func (s *InviteServer) DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error {
	return s.inviteRepo.DeclineInvite(ctx, userID, roomID)
}
