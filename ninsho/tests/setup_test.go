package tests

import (
	"context"
	"fmt"
	"ninsho/internal/gen/auth"
	"ninsho/internal/server"
	"testing"
	"time"

	"google.golang.org/grpc"
)

var conn auth.AuthRoutesClient

func init() {

	config := server.CreateDefaultConfig()

	server_conn := server.NewServer(config, testDB)
	go server_conn.Serve()

	full_address := config.ToString()

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}

	ctx, close := context.WithTimeout(context.Background(), time.Second)
	defer close()

	if normal_conn, err := grpc.DialContext(ctx, full_address, opts...); err == nil {
		conn = auth.NewAuthRoutesClient(normal_conn)
	} else {
		panic(fmt.Sprintf("Error while trying to connect to address '%s'", full_address))
	}
}

func TestValidation(t *testing.T) {
	t.Run("Testing registers", toTestRegister)
	t.Run("Testing login", toTestLogin)
}
