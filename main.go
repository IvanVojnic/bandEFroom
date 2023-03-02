package main

import (
	"github.com/IvanVojnic/bandEFroom/internal/service"
	"net"

	"github.com/IvanVojnic/bandEFroom/internal/config"
	"github.com/IvanVojnic/bandEFroom/internal/repository"
	"github.com/IvanVojnic/bandEFroom/internal/rpc"
	pr "github.com/IvanVojnic/bandEFroom/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
		logrus.WithFields(logrus.Fields{
			"Error connection to database rep.NewPostgresDB()": err,
		}).Fatal("DB ERROR CONNECTION")
	}
	defer repository.ClosePool(db)
	inviteRepo := repository.NewRoomPostgres(db)
	roomRepo := repository.NewRoomPostgres(db)
	inviteServ := service.NewInviteServer(inviteRepo)
	roomServ := service.NewRoomServer(roomRepo)
	inviteGRPC := rpc.NewInviteServer(inviteServ)
	roomGRPC := rpc.NewRoomServer(roomServ)

	pr.RegisterRoomServer(s, inviteGRPC)
	pr.RegisterRoomServer(s, roomGRPC)
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		defer logrus.Fatalf("error while listening port: %e", err)
	}

	if errServ := s.Serve(listen); errServ != nil {
		defer logrus.Fatalf("error while listening server: %e", err)
	}
}
