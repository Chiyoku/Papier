package factory

import (
	"shopping/internal/repositories"
	"shopping/internal/usecase"

	"gorm.io/gorm"
)

func CommentRepoFactory(db *gorm.DB) repositories.CommentRepo {
	return &repositories.CommentRepoImpl{
		Adapter: db,
	}
}

func CommentUseCaseFactory(commentRepository *repositories.CommentRepo) *usecase.CommentUseCase {
	return &usecase.CommentUseCase{}
}
