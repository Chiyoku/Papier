package tests

import (
	"ninsho/internal/gen/auth"
	"ninsho/internal/users"
	"testing"
)

func TestValidation(t *testing.T) {
	jwt := testValidLogin(t, &auth.LoginRequest{
		Email:    "eta@hotmail.com",
		Password: "lerolero123",
	})
	if err := users.ValidateJWT(jwt); err != nil {
		t.Fatalf("Failed to verify the JWT for %s", "eta@hotmail.com")
	}
}
