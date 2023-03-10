package repository

import (
	"context"
	"fmt"
	"github.com/IvanVojnic/bandEFroom/models"

	prNotif "github.com/IvanVojnic/bandEFnotif/proto"

	"github.com/google/uuid"
	"time"
)

// NotificationMS is a wrapper to ms object
type NotificationMS struct {
	client prNotif.InviteRoomClient
}

// NewNotificationMS used to init NotificationMS
func NewNotificationMS(client prNotif.InviteRoomClient) *NotificationMS {
	return &NotificationMS{client: client}
}

func (r *NotificationMS) StorageInvite(ctx context.Context, userCreator models.User, usersInvited []*models.User, roomID uuid.UUID, date time.Time, place string) error {
	userCreatorGRPC := &prNotif.User{UserID: userCreator.ID.String(), UserName: userCreator.Name, UserEmail: userCreator.Email}
	var usersInvitedGRPC []*prNotif.User
	for _, userInvited := range usersInvited {
		usersInvitedGRPC = append(usersInvitedGRPC, &prNotif.User{UserID: userInvited.ID.String(), UserName: userInvited.Name, UserEmail: userInvited.Email})
	}
	_, errGRPC := r.client.StorageInvite(ctx, &prNotif.StorageInviteRequest{UserCreator: userCreatorGRPC, Users: usersInvitedGRPC, RoomID: roomID.String(), Place: place, Date: date.String()})
	if errGRPC != nil {
		return fmt.Errorf("error while storage notiffications of invite, %s", errGRPC)
	}
	return nil
}
