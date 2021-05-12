package main

import (
	"context"
	"fmt"
	"ninsho/internal/gen/auth"
	"ninsho/internal/server"
	"time"

	"google.golang.org/grpc"
)

func main() {

	config := server.CreateDefaultConfig()
	full_address := config.ToString()

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}

	ctx, close := context.WithTimeout(context.Background(), time.Second)
	defer close()

	if normal_conn, err := grpc.DialContext(ctx, full_address, opts...); err == nil {
		auth := auth.NewAuthRoutesClient(normal_conn)
		lero(auth)
	} else {
		panic(fmt.Sprintf("Error while trying to connect to address '%s'", full_address))
	}
}

func lero(conn auth.AuthRoutesClient) {
	user := &auth.LoginRequest{
		Email:    "invalidemail@gmail.com",
		Password: "lerolero123",
	}

	jwt, err := conn.Login(context.Background(), user)
	if err != nil {
		fmt.Printf("Expected a sucessful login while trying to login with email '%s'\n", err)
	} else {

		fmt.Printf("SUCESS: %s\n", jwt)
	}
}
