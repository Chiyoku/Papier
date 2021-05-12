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

func testValidLogin(t *testing.T, user *auth.LoginRequest) *auth.Response {
	jwt, err := conn.Login(context.Background(), user)
	if err != nil {
		t.Errorf("Expected a sucessful login while trying to login with email '%s'", user.Email)
		panic(err)
	}
	return jwt
}

func testInvalidLogin(t *testing.T, user *auth.LoginRequest) *auth.Response {
	response, err := conn.Login(context.Background(), user)
	if err != nil {
		t.Errorf("Expected an error while trying to login with '%s'", user.Email)
		panic(err)
	}
	return response
}

func Test(t *testing.T) {
	t.Run("Testing registers", ToTestRegister)
	t.Run("Testing login", ToTestLogin)
}

func ToTestLogin(t *testing.T) {
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
