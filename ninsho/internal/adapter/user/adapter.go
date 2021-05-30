package user

import "ninsho/internal/models"

type UserAdapter interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(id int) (*models.User, error)
	GetUserByEmail(nickname string) (*models.User, error)
}
