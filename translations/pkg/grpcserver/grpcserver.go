package grpcserver

import (
	"log"
	"net"
	"time"
	"translations/config"
	grpcrouter "translations/internal/presentation/grpc"

	"google.golang.org/grpc"
)

type grpcserver struct {
	Server *grpc.Server
}

func Init(cfg config.GRPCServerConfig, router *grpcrouter.GRPCRouter) *grpcserver {
	var server *grpc.Server

	server = grpc.NewServer(grpc.ConnectionTimeout(time.Second * 5))

	router.Register(server)

	listener, err := net.Listen("tcp", cfg.Address+":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go server.Serve(listener)

	return &grpcserver{
		Server: server,
	}
}

func (gs *grpcserver) Shutdown() {
	gs.Server.GracefulStop()
}
