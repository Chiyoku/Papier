package server

import (
	"context"
	"errors"
	"ninsho/internal/gen/auth"
	"ninsho/internal/jwt"
	"ninsho/internal/models"
)

func (s Server) Login(ctx context.Context, body *auth.LoginRequest) (*auth.Response, error) {
	user, err := s.userService.Login(body.Email, body.Password)
	if err != nil {
		return nil, errors.New("not found")
	}

	jwt, err := jwt.GenerateJWT([]byte("teste"), user)

	if err != nil {
		return nil, err
	}

	return &auth.Response{JWT: jwt}, nil
}

func (s Server) Register(ctx context.Context, body *auth.RegisterRequest) (*auth.Response, error) {

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
