package tests

import (
	"context"
	"ninsho/internal/gen/auth"
	"testing"
)

func testValidLogin(t *testing.T, user *auth.LoginRequest) *auth.Response {
	jwt, err := conn.Login(context.Background(), user)
	if err != nil {
		t.Errorf("Expected a sucessful login while trying to login with email '%s'", user.Email)
	}
	return jwt
}

func testInvalidLogin(t *testing.T, user *auth.LoginRequest) *auth.Response {
	response, err := conn.Login(context.Background(), user)
	if err == nil {
		t.Errorf("Expected an error while trying to login with '%s'", user.Email)
	}
	return response
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
