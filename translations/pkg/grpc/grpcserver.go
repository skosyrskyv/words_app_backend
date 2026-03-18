package grpcserver

import (
	"log"
	"net"
	"time"
	"translations/config"

	"google.golang.org/grpc"
)

type Server struct {
	s *grpc.Server
}

func Init(cfg config.GRPCServerConfig) (*Server, error) {
	var server *grpc.Server

	server = grpc.NewServer(grpc.ConnectionTimeout(time.Second * 5))

	listener, err := net.Listen("tcp", cfg.Address+":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server.Serve(listener)

	return &Server{
		s: server,
	}, nil
}
