package user

import (
	"errors"
	"ninsho/internal/adapter/user"
	"ninsho/internal/models"
)

type UserService struct {
	params  *hashParams
	adapter user.UserAdapter
}

func NewUserService(params *hashParams, adapter user.UserAdapter) *UserService {
	return &UserService{params, adapter}
}

func (service *UserService) Login(email string, password string) (*models.User, error) {
	user, err := service.adapter.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	res, err := Verify(service.params, user.PasswordHash, password)
	if err != nil {
		return nil, err
	}
	if res {
		return user, nil
	}
	return nil, errors.New("passwords not match")
}

func (service *UserService) Register(user *models.User) (*models.User, error) {
	user, err := service.adapter.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) Hash(password string) (string, error) {
	res, err := Hash(service.params, password)
	return string(res), err
}

func ValidateJWT(jwt string) bool {
	return false
}
