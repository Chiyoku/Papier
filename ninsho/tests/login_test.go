package tests

import (
	"context"
	"flag"
	"fmt"
	"ninsho/internal/gen/auth"
	"ninsho/internal/server"
	"testing"

	"google.golang.org/grpc"
)

var conn auth.AuthRoutesClient

func init() {
	address := flag.String("addr", "", "address of the server")
	port := flag.Int("port", 8080, "port of the server")

	server_conn := server.NewServer(address, port, testDB)
	defer server_conn.Close()
	go server_conn.Serve()

	full_address := fmt.Sprintf("%s:%d", address, port)
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	if normal_conn, err := grpc.Dial(full_address, opts...); err == nil {
		conn = auth.NewAuthRoutesClient(normal_conn)
	} else {
		panic(fmt.Sprintf("Error while trying to connect to %s", full_address))
	}
}

func testValidLogin(t *testing.T, user *auth.LoginRequest) *auth.Response {
	jwt, err := conn.Login(context.Background(), user)
	if err != nil {
		t.Errorf("Expected a sucessful login while trying to login with email '%s'", user.Email)
	}
	return jwt
}

func testInvalidLogin(t *testing.T, user *auth.LoginRequest) *auth.Response {
	response, err := conn.Login(context.Background(), user)
	if err != nil {
		t.Errorf("Expected an error while trying to login with '%s'", user.Email)
	}
	return response
}

func TestLogin(t *testing.T) {
	testValidLogin(t, &auth.LoginRequest{
		Email:    "eta@hotmail.com",
		Password: "lerolero123",
	})

	testInvalidLogin(t, &auth.LoginRequest{
		Email:    "notexists@hotmail.com",
		Password: "lerolero123",
	})

	testInvalidLogin(t, &auth.LoginRequest{
		Email:    "minimumsize@hotmail.com",
		Password: "lerole",
	})

	testInvalidLogin(t, &auth.LoginRequest{
		Email:    "invalidemail.com",
		Password: "lerolero123",
	})

	testInvalidLogin(t, &auth.LoginRequest{
		Email:    "invalidemail.com",
		Password: "lerolero123",
	})
}
