package user

import (
	"ninsho/internal/models"

	"gorm.io/gorm"
)

type GormUserAdapter struct {
	DB *gorm.DB
}

func (adapter *GormUserAdapter) CreateUser(user *models.User) (int, error) {
	result := adapter.DB.Create(user)
	return user.ID, result.Error
}

func (adapter *GormUserAdapter) GetUser(id int) (*models.User, error) {
	user := &models.User{ID: id}
	result := adapter.DB.First(user)
	return user, result.Error
}
