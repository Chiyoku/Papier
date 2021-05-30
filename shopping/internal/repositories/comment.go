package repositories

import (
	"shopping/internal/model"

	"gorm.io/gorm"
)

type CommentRepo interface {
	CreateComment(comment *model.Comment) error
	FindComment(id string) (*model.Comment, error)
	DeleteComment(id string) error
	GetAll(id string) ([]model.Comment, error)
}

type CommentRepoImpl struct {
	Adapter *gorm.DB
}

func (repo *CommentRepoImpl) CreateComment(comment *model.Comment) error {
	err := repo.Adapter.Create(comment).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *CommentRepoImpl) FindComment(id string) (*model.Comment, error) {
	var comment model.Comment

	err := repo.Adapter.First(&comment, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}


func (repo *CommentRepoImpl) DeleteComment(id string) error {
	err := repo.Adapter.Delete(&model.Comment{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *CommentRepoImpl) GetAll(id string) ([]model.Comment, error) {

	return nil, nil
}
