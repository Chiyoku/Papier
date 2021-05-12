package user

import "ninsho/internal/models"

type UserServiceImpl interface {
	Login(username string, password string) (*models.User, error)
	Register(user *models.User) (*models.User, error)
	Hash(password string) (string, error)
}
