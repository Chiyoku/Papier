package usecase

import (
	"shopping/internal/model"
	"shopping/internal/repositories"
)

type CommentUseCase struct {
	CommentRepo repositories.CommentRepo
}

func (commentUseCase *CommentUseCase) Create(message, author, video string) (*model.Comment, error) {
	return nil, nil
}
