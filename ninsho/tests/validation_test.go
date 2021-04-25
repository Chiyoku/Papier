package tests

import (
	"ninsho/internal/gen/auth"
	"ninsho/internal/users"
	"testing"
)

func TestValidation(t *testing.T) {
	res := testValidLogin(t, &auth.LoginRequest{
		Email:    "eta@hotmail.com",
		Password: "lerolero123",
	})
	if !users.ValidateJWT(res.JWT) {
		t.Fatalf("Failed to verify the JWT for %s", "eta@hotmail.com")
	}
}
