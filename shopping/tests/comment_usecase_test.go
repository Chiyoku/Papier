package tests

import (
	"shopping/internal/model"
	"shopping/internal/usecase"
	"testing"
)

func testValidGetCommentUseCase(t *testing.T, usecase *usecase.CommentUseCase) {
	comment, err := usecase.Create("dkapowdkaopskdlçkzzz", "kdkopkddkkddkkdkdkd", "awopkapkpoqqqq")

	if err != nil {
		t.Errorf("Not expected an error to create a comment to try to get after")
	}
}

func testInvalidGetCommentUseCase(t *testing.T, usecase *usecase.CommentUseCase) {
	_, err := usecase.GetByID("doakwpdkawdp")
	
	if err != nil {
		t.Errorf("Expected an error to try to get a comment that doesn't exists in the database")
	}
}

func testCommentUseCaseCreate(t *testing.T, usecase *usecase.CommentUseCase) *model.Comment {
	comment, err := usecase.Create("oi", "ppppp", "ppp")

	if err != nil {
		t.Errorf("Not expected an error to create the comment(usecase)")
	}

	return comment
}

func testCommentUseCase(t *testing.T, usecase *usecase.CommentUseCase) {
	testCommentUseCaseCreate(t, usecase)
	
	testValidGetCommentUseCase(t, usecase)
	testInvalidGetCommentUseCase(t, usecase)
}
