package server

import (
	"context"
	"fmt"
	"net"
	"ninsho/internal/gen/auth"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	db       *gorm.DB
	server   *grpc.Server
	listener *net.Listener
}

func (s Server) Login(ctx context.Context, body *auth.LoginRequest) (*auth.Response, error) {
	return nil, nil
}

func (s Server) Register(ctx context.Context, body *auth.RegisterRequest) (*auth.Response, error) {
	return nil, nil
}

func (s Server) Validate(ctx context.Context, body *auth.ValidationRequest) (*auth.ValidationResponse, error) {
	return nil, nil
}

func (conn *Server) Close() {
	conn.server.Stop()
}

func (conn *Server) Serve() {
	conn.server.Serve(*conn.listener)
}

func NewServer(address string, port string, db *gorm.DB) *Server {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	conn := Server{
		db:       db,
		server:   grpcServer,
		listener: &lis,
	}
	auth.RegisterAuthRoutesServer(grpcServer, conn)
	return &conn
}
