package tests

import (
	"shopping/internal/model"
	"shopping/internal/repositories"
	"testing"
)

func testCreateComment(t *testing.T, repo repositories.CommentRepo) {
	comment, err := model.NewComment("hello world", "ddd", "bbb")

	if err != nil {
		t.Errorf("Not expected an error to create the comment")
	}

	repoErr := repo.CreateComment(comment)

	if repoErr != nil {
		t.Errorf("Not expected an error to save the comment in the database")
	}
}

func testValidFindComment(t *testing.T, repo repositories.CommentRepo) {
	comment, err := model.NewComment("hello world", "ddd", "bbb")

	if err != nil {
		t.Errorf("Not expected an error to create the comment")
	}

	createErr := repo.CreateComment(comment)

	if createErr != nil {
		t.Errorf("Not expected an error to save the comment in the database")
	}

	_, findErr := repo.FindComment(comment.ID)

	if findErr != nil {
		t.Errorf("Not expected an error to find a comment that actually exists in the database ;-;")
	}
}

func testInvalidFindComment(t *testing.T, repo repositories.CommentRepo) {
	_, err := repo.FindComment("oooo")

	if err == nil {
		t.Errorf("Expected an error to try to get a comment that doesn't exists in the database")
	}
}

func testCommentRepo(t *testing.T, repo repositories.CommentRepo) {
	testCreateComment(t, repo)
	testValidFindComment(t, repo)
	testInvalidFindComment(t, repo)
}