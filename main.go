// Package main is a main package used to start the program
package main

import (
	"net"
	"os"

	prNotif "github.com/IvanVojnic/bandEFnotif/proto"
	"github.com/IvanVojnic/bandEFroom/internal/config"
	"github.com/IvanVojnic/bandEFroom/internal/repository"
	"github.com/IvanVojnic/bandEFroom/internal/rpc"
	"github.com/IvanVojnic/bandEFroom/internal/service"
	pr "github.com/IvanVojnic/bandEFroom/proto"
	prUser "github.com/IvanVojnic/bandEFuser/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	cfg, err := config.NewConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error":  err,
			"config": cfg,
		}).Fatal("failed to get config")
	}
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("DB ERROR CONNECTION %s", err)
	}
	defer repository.ClosePool(db)

	connUserMS, err := grpc.Dial(os.Getenv("PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("error while conecting to user ms, %s", err)
	}
	clientUserComm := prUser.NewUserCommClient(connUserMS)

	connNotifMS, err := grpc.Dial(os.Getenv("PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("error while conecting to notif ms, %s", err)
	}
	clientNotifComm := prNotif.NewInviteRoomClient(connNotifMS)

	inviteRepo := repository.NewInvitePostgres(db)
	roomRepo := repository.NewRoomPostgres(db)
	userRepo := repository.NewUserMS(clientUserComm)
	notifRepo := repository.NewNotificationMS(clientNotifComm)

	inviteServ := service.NewInviteServer(inviteRepo, userRepo, notifRepo)
	roomServ := service.NewRoomServer(roomRepo)

	inviteGRPC := rpc.NewInviteServer(inviteServ)
	roomGRPC := rpc.NewRoomServer(roomServ)

	pr.RegisterInviteServer(s, inviteGRPC)
	pr.RegisterRoomServer(s, roomGRPC)
	listen, err := net.Listen("tcp", "0.0.0.0:8000") // ???????????? ???????????? ????????????
	if err != nil {
		defer logrus.Errorf("error while listening port: %e", err)
	}
	if errServ := s.Serve(listen); errServ != nil {
		defer logrus.Errorf("error while listening server: %e", err)
	}
}
