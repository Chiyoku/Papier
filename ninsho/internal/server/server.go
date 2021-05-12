package server

import (
	"net"
	adapter "ninsho/internal/adapter/user"
	"ninsho/internal/gen/auth"
	"ninsho/internal/user"

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
	userService user.UserServiceImpl
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

	params := user.NewHashParams()

	service := user.NewUserService(params, adapter.NewGormAdapter(db))

	conn := Server{
		userService: service,
		server:      grpcServer,
		listener:    &listener,
		listening:   false,
	}

	auth.RegisterAuthRoutesServer(grpcServer, conn)

	return &conn
}
