package grpcserver

import (
	"log"
	"net"
	"time"
	"translations/config"

	"google.golang.org/grpc"
)

type grpcserver struct {
	s *grpc.Server
}

func Init(cfg config.GRPCServerConfig) (*grpcserver, error) {
	var server *grpc.Server

	server = grpc.NewServer(grpc.ConnectionTimeout(time.Second * 5))

	listener, err := net.Listen("tcp", cfg.Address+":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go server.Serve(listener)

	return &grpcserver{
		s: server,
	}, nil
}

func (gs *grpcserver) Shutdown() {
	gs.s.GracefulStop()
}
