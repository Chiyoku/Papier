package user

import "ninsho/internal/user"

type UserAdapter interface {
	CreateUser(user *user.User) (int, error)
	GetUser(id int) (*user.User, error)
}
