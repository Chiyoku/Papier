package tests

import (
	"context"
	"ninsho/internal/gen/auth"
	"testing"
)

func testValidRegister(t *testing.T, user *auth.RegisterRequest) {
	_, err := conn.Register(context.Background(), user)
	if err != nil {
		t.Errorf("Expected a sucessful Register while trying to Register with email '%s': %s", user.Email, err)
	} else {
		t.Logf("Sucessfully registered '%s'", user.Email)
	}
}

func testInvalidRegister(t *testing.T, user *auth.RegisterRequest) *auth.Response {
	response, err := conn.Register(context.Background(), user)
	if err == nil {
		t.Errorf("Expected an error while trying to Register with '%s' username '%s'", user.Email, user.Username)
	}
	return response
}

func toTestRegister(t *testing.T) {
	testValidRegister(t, &auth.RegisterRequest{
		Username: "chiyoku10",
		Email:    "notexistsal@hotmail.com",
		Password: "lerolero123",
	})

	testValidRegister(t, &auth.RegisterRequest{
		Username: "chiyoku1",
		Email:    "eta@hotmail.com",
		Password: "lerolero123",
	})

	testValidRegister(t, &auth.RegisterRequest{
		Username: "chiyoku",
		Email:    "notexistsbutwillexistsnow@hotmail.com",
		Password: "lerolero123",
	})

	testInvalidRegister(t, &auth.RegisterRequest{
		Username: "chiyoku1",
		Email:    "eta@hotmail.com",
		Password: "lerolero123",
	})

	testInvalidRegister(t, &auth.RegisterRequest{
		Username: "chiyoku2",
		Email:    "minimumsize@hotmail.com",
		Password: "lerole",
	})

	testInvalidRegister(t, &auth.RegisterRequest{
		Username: "chiyoku3",
		Email:    "invalidemail.com",
		Password: "lerolero123",
	})

	testInvalidRegister(t, &auth.RegisterRequest{
		Username: "chiyoku4",
		Email:    "invalidemail.com",
		Password: "lerolero123",
	})

}
