// Package repository define room repo methods
package repository

import (
	"context"
	"fmt"

	"github.com/IvanVojnic/bandEFroom/models"
	pr "github.com/IvanVojnic/bandEFuser/proto"

	"github.com/google/uuid"
)

// UserPostgres is a wrapper to db object
type UserPostgres struct {
	client pr.UserCommClient
}

// NewUserPostgres used to init UserPostgres
func NewUserPostgres(client pr.UserCommClient) *UserPostgres {
	return &UserPostgres{client: client}
}

// GetUsers used to get full info of users from userMS
func (r *UserPostgres) GetUsers(ctx context.Context, usersID []*uuid.UUID) ([]*models.User, error) {
	users := make([]*models.User, 0)
	usersIDStr := make([]string, 0)
	for _, ID := range usersID {
		usersIDStr = append(usersIDStr, ID.String())
	}

	res, errGRPC := r.client.GetUsers(ctx, &pr.GetUsersRequest{UsersID: usersIDStr})
	if errGRPC != nil {
		return users, fmt.Errorf("error while sign up, %s", errGRPC)
	}

	for _, user := range res.Users {
		userID, errUserID := uuid.Parse(user.ID)
		if errUserID != nil {
			return users, fmt.Errorf("error while parsing room ID, %s", errUserID)
		}
		users = append(users, &models.User{ID: userID, Name: user.Name, Email: user.Email})
	}
	return users, nil
}

// GetUser used to get full info of users from userMS
func (r *UserPostgres) GetUser(ctx context.Context, userID uuid.UUID) (models.User, error) {
	res, errGRPC := r.client.GetUser(ctx, &pr.GetUserRequest{UserID: userID.String()})
	if errGRPC != nil {
		return models.User{}, fmt.Errorf("error while sign up, %s", errGRPC)
	}
	userID, errUserID := uuid.Parse(res.ID)
	if errUserID != nil {
		return models.User{}, fmt.Errorf("error while parsing room ID, %s", errUserID)
	}
	user := models.User{ID: userID, Name: res.Name, Email: res.Email}
	return user, nil
}
