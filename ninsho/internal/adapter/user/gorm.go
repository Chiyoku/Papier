package user

import (
	"ninsho/internal/models"

	"gorm.io/gorm"
)

type GormUserAdapter struct {
	DB *gorm.DB
}

func NewGormAdapter(db *gorm.DB) *GormUserAdapter {
	return &GormUserAdapter{
		DB: db,
	}
}

func (adapter *GormUserAdapter) CreateUser(user *models.User) (*models.User, error) {
	result := adapter.DB.Create(user)
	return user, result.Error
}

func (adapter *GormUserAdapter) GetUser(id int) (*models.User, error) {
	user := &models.User{ID: id}
	result := adapter.DB.First(user, id)
	return user, result.Error
}

func (adapter *GormUserAdapter) GetUserByUsername(nickname string) (*models.User, error) {
	user := &models.User{}
	result := adapter.DB.First(user).Where("email = ?", nickname)
	return user, result.Error
}
