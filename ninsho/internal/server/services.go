package server

import (
	"context"
	"errors"
	"ninsho/internal/gen/auth"
)

func (s Server) Login(ctx context.Context, body *auth.LoginRequest) (*auth.Response, error) {
	return nil, errors.New("login not implemented yet")
}

func (s Server) Register(ctx context.Context, body *auth.RegisterRequest) (*auth.Response, error) {
	return nil, errors.New("register not implemented yet")
}

func (s Server) Validate(ctx context.Context, body *auth.ValidationRequest) (*auth.ValidationResponse, error) {
	return nil, errors.New("validation not implemented yet")
}
