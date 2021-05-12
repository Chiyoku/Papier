package server

import (
	"fmt"
	"net"
	"ninsho/internal/gen/auth"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	db        *gorm.DB
	server    *grpc.Server
	listener  *net.Listener
	listening bool
}

func (conn *Server) Close() {
	conn.server.Stop()
}

func (conn *Server) Serve() {
	conn.server.Serve(*conn.listener)
}

func NewServer(config *ServerConfig, db *gorm.DB) *Server {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Address, config.Port))

	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	conn := Server{
		db:        db,
		server:    grpcServer,
		listener:  &lis,
		listening: false,
	}

	auth.RegisterAuthRoutesServer(grpcServer, conn)
	return &conn
}
