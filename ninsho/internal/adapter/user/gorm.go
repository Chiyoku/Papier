package user

import (
	"ninsho/internal/user"

	"gorm.io/gorm"
)

type GormUserAdapter struct {
	DB *gorm.DB
}

func (adapter *GormUserAdapter) CreateUser(user *user.User) (int, error) {
	result := adapter.DB.Create(user)
	return user.ID, result.Error
}

func (adapter *GormUserAdapter) GetUser(id int) (*user.User, error) {
	user := &user.User{ID: id}
	result := adapter.DB.First(user)
	return user, result.Error
}
