package tests

import (
	"context"
	"fmt"
	"ninsho/internal/gen/auth"
	"ninsho/internal/server"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
)

var conn auth.AuthRoutesClient

func init() {
	address := os.Getenv("ADDR")
	port := os.Getenv("PORT")

	server_conn := server.NewServer(address, port, testDB)
	defer server_conn.Close()
	go server_conn.Serve()

	full_address := fmt.Sprintf("%s:%s", address, port)
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
