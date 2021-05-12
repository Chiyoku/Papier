package jwt

import (
	"ninsho/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(key []byte, user *models.User) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	// Just to test
	atClaims["email"] = user.Email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString(key)
}
