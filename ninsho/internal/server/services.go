package server

import (
	"context"
	"errors"
	"ninsho/internal/gen/auth"
	"ninsho/internal/jwt"
	"ninsho/internal/models"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (s Server) Login(ctx context.Context, body *auth.LoginRequest) (*auth.Response, error) {
	user, err := s.userService.Login(body.Email, body.Password)
	if err != nil {
		return nil, err
	}

	jwt, err := jwt.GenerateJWT([]byte("teste"), user)

	if err != nil {
		return nil, err
	}

	return &auth.Response{JWT: jwt}, nil
}

func (s Server) Register(ctx context.Context, body *auth.RegisterRequest) (*auth.Response, error) {
	err := isValidPassword(body.Password)

	if err != nil {
		return nil, err
	}

	err = isValidEmail(body.Email)

	if err != nil {
		return nil, err
	}

	hash, err := s.userService.Hash(body.Password)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: hash,
	}

	user, err = s.userService.Register(user)

	if err != nil {
		return nil, err
	}

	jwt, err := jwt.GenerateJWT([]byte("teste"), user)

	if err != nil {
		return nil, err
	}

	return &auth.Response{JWT: jwt}, nil
}

func (s Server) Validate(ctx context.Context, body *auth.ValidationRequest) (*auth.ValidationResponse, error) {
	return nil, errors.New("validation not implemented yet")
}

func isValidPassword(pass string) error {
	if len(pass) < 8 || len(pass) >= 256 {
		return errors.New("the password is invalid in the range of 8 to 256")
	}
	return nil
}

func isValidEmail(email string) error {
	if len(email) < 3 && len(email) > 254 {
		return errors.New("invalid length of email")
	}
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email")
	}
	return nil
}
