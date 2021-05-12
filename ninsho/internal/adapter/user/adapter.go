package user

import "ninsho/internal/models"

type UserAdapter interface {
	CreateUser(user *models.User) (int, error)
	GetUser(id int) (*models.User, error)
}
