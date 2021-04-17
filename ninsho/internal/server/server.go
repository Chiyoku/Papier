package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"ninsho/internal/gen/auth"

	"github.com/go-pg/pg"
	"google.golang.org/grpc"
)

type Server struct {
	db *pg.DB
}

func (s Server) Login(ctx context.Context, body *auth.LoginRequest) (*auth.Response, error) {
	return nil, nil
}

func (s Server) Register(ctx context.Context, body *auth.RegisterRequest) (*auth.Response, error) {
	return nil, nil
}

func (s Server) Validate(ctx context.Context, body *auth.ValidationRequest) (*auth.Response, error) {
	return nil, nil
}

func NewServer(address string, port int, db *pg.DB) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	auth.RegisterAuthRoutesServer(grpcServer, Server{db: db})
	grpcServer.Serve(lis)
	return grpcServer
}
