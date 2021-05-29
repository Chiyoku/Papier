package jwt

import (
	"errors"
	"fmt"
	"ninsho/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaim struct {
	Id          int
	Permissions uint32
	jwt.StandardClaims
}

func GenerateJWT(key []byte, user *models.User) (string, error) {
	atClaims := &JwtClaim{
		Id:          user.ID,
		Permissions: 0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 10).Unix(),
		},
	}

	jwtInst := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return jwtInst.SignedString(key)
}

func ValidateJWT(key []byte, jwtToken string) (*JwtClaim, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaim); ok {
		if claims.ExpiresAt < time.Now().Local().Unix() {
			return nil, errors.New("expired token")
		}
		return claims, nil
	}

	return nil, errors.New("couldnt parse claims")

}
