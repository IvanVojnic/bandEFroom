package service

import (
	"context"
	"fmt"
	"time"

	"github.com/IvanVojnic/bandEFroom/models"

	"github.com/google/uuid"
)

// Invite interface consists of methods to invite users
type Invite interface {
	SendInvite(ctx context.Context, users *[]models.User, roomID uuid.UUID, creatorID uuid.UUID) error
	AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error
	DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error
	CreateRoom(ctx context.Context, userCreatorID uuid.UUID, place string, date time.Time) (uuid.UUID, error)
}

// InviteServer define service invites
type InviteServer struct {
	inviteRepo Invite
}

// NewInviteServer used to init service user communicate struct
func NewInviteServer(inviteRepo Invite) *InviteServer {
	return &InviteServer{inviteRepo: inviteRepo}
}

// SendInvite used to invite friends to the room by repo
func (s *InviteServer) SendInvite(ctx context.Context, userCreatorID uuid.UUID, users *[]models.User, place string, date time.Time) error {
	roomID, err := s.inviteRepo.CreateRoom(ctx, userCreatorID, place, date)
	if err != nil {
		return fmt.Errorf("error while creating room, %s", err)
	}
	return s.inviteRepo.SendInvite(ctx, users, roomID, userCreatorID)
}

// AcceptInvite used to accept invite to the room by repo
func (s *InviteServer) AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {
	return s.inviteRepo.AcceptInvite(ctx, userID, roomID)
}

// DeclineInvite used to accept invite to the room by repo
func (s *InviteServer) DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {
	return s.inviteRepo.DeclineInvite(ctx, userID, roomID)
}
