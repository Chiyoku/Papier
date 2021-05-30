package tests

import (
	"context"
	"ninsho/internal/gen/auth"
	"testing"
)

func testValidLogin(t *testing.T, user *auth.LoginRequest) {
	_, err := conn.Login(context.Background(), user)
	if err != nil {
		t.Errorf("Expected a sucessful login while trying to login with email '%s' but got '%s'", user.Email, err)
	}
}

func testInvalidLogin(t *testing.T, user *auth.LoginRequest) {
	_, err := conn.Login(context.Background(), user)
	if err == nil {
		t.Errorf("Expected an error while trying to login with '%s'", user.Email)
	}
}

func toTestLogin(t *testing.T) {

	testValidLogin(t, &auth.LoginRequest{
		Email:    "notexistsal@hotmail.com",
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
