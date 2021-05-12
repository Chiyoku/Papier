package server

import (
	"net"
	"ninsho/internal/adapter/user"
	"ninsho/internal/gen/auth"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

/* TODO: Probably Server is not a good name
 * because it unfortunately has a instance of the database
 * inside each of adapters
 * in order to avoid the usage of a global variable inside
 * the services in services.go
 */
type Server struct {
	userAdapter user.UserAdapter
	server      *grpc.Server
	listener    *net.Listener
	listening   bool
}

func (conn *Server) Close() {
	conn.server.Stop()
}

func (conn *Server) Serve() {
	conn.server.Serve(*conn.listener)
}

func NewServer(config *ServerConfig, db *gorm.DB) *Server {
	listener, err := net.Listen("tcp", config.ToString())

	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	conn := Server{
		userAdapter: &user.GormUserAdapter{
			DB: db,
		},
		server:    grpcServer,
		listener:  &listener,
		listening: false,
	}

	auth.RegisterAuthRoutesServer(grpcServer, conn)

	return &conn
}
