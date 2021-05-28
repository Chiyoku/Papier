package __tests__

import (
	"os"
	
	"shopping/internals/factory"
	"shopping/internals/usecase"
	"shopping/services"

	"shopping/logic/model"
	"testing"

	"gorm.io/gorm"
)

func testInvalidComment(t *testing.T, usecase *usecase.CommentUseCase) {
	_, err := usecase.Get("oooooo")

	if err == nil {
		t.Errorf("Expected an error while trying to get a commen that doesn't exist")
	}
}

func testValidComment(t *testing.T, usecase *usecase.CommentUseCase) {
	comment, err := usecase.Create("message", "ddd11", "ddd111")

	if err != nil {
		t.Errorf("Not expected an error to create a valid comment")
	}

	_, err := usecase.Get(comment.ID)

	if err != nil {
		t.Errorf("Not expected an error to get a comment cause it actually exists in the database lol")
	}
}

func testCommentUseCase(t *testing.T, usecase *usecase.CommentUseCase) {
	comment, err := usecase.Create("hello world", "999111", "222333")

	if err != nil {
		t.Errorf("Nott expected an error to create the comment")
	}
}

func testNewComment(t *testing.T) {
	comment, err := model.NewComment("hello world", "aa", "bb")

	if err != nil {
		t.Errorf("Not expected an error to create the comment")
	}
	
	if comment.ID == "" {
		t.Errorf("Expected an uuid as id in the comment, but got an empty string :(")
	}
}

func testComment(t *testing.T, database *gorm.DB) {
	usecase, err := factory.NewCommentUseCase(database)

	if err != nil {
		t.Errorf("Not expected an error to create the comment usecase")
	}

	t.Run("Testing comment domain model", testNewComment)
	t.Run("Testing comment usecase", func(t *testing.T) {
		testValidComment(t, usecase)
		testInvalidComment(t, usecase)
	})
}