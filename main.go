package main

import (
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
	//userCommRepo := repository.NewUserCommPostgres(db)
	inviteServ := service.NewUserAuthServer(inviteRepo)
	//userCommServ := service.NewUserCommServer(userCommRepo)
	inviteGRPC := rpc.NewInviteServer(inviteServ)
	//userCommGRPC := rpc.NewUserCommServer(userCommServ)

	//pr.RegisterUserServer(s, userAuthGRPC)
	pr.RegisterRoomServer(s, inviteGRPC)
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		defer logrus.Fatalf("error while listening port: %e", err)
	}

	if errServ := s.Serve(listen); errServ != nil {
		defer logrus.Fatalf("error while listening server: %e", err)
	}
}
