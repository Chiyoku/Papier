package tests

import (
	"context"
	"ninsho/internal/gen/auth"
	"ninsho/internal/jwt"
	"testing"
)

func toTestJwt(t *testing.T) {
	response, err := conn.Login(context.Background(), &auth.LoginRequest{
		Email:    "notexistsal@hotmail.com",
		Password: "lerolero123",
	})
	if err != nil {
		t.Errorf("Expected a sucessful login while trying to login with email '%s' but got '%s'", "notexistsal@hotmail.com", err)
	}
	_, err = jwt.ValidateJWT(([]byte)("keyDeTeste"), response.JWT)
	if err != nil {
		t.Errorf("Error while trying to decode the jwt '%s'", err)
	}
}
