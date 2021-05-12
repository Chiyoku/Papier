package tests

import (
	"context"
	"ninsho/internal/gen/auth"
	"testing"
)

func testValidRegister(t *testing.T, user *auth.RegisterRequest) {
	_, err := conn.Register(context.Background(), user)
	if err != nil {
		t.Errorf("Expected a sucessful Register while trying to Register with email '%s'", user.Email)
	}
}

func testInvalidRegister(t *testing.T, user *auth.RegisterRequest) *auth.Response {
	response, err := conn.Register(context.Background(), user)
	if err == nil {
		t.Errorf("Expected an error while trying to Register with '%s'", user.Email)
	}
	return response
}
