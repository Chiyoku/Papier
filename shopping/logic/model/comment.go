package model

import (
	"time"

	"github.com/myesui/uuid"
)

type Comment struct {
	Base
	Message  string `json:"message"`
	AuthorID string `json:"author_id"`
	VideoID  string `json:"video_id"`
}

type CommentRepo interface {
	Add(comment *Comment) error
	GetByID(id string) (*Comment, error)
	GetByVideo(videoID string) ([]Comment, error)
}

func NewComment(message, authorID, videoID string) (*Comment, error) {
	comment := Comment{
		Message: message,
		AuthorID: authorID,
		VideoID: videoID,

	}

	comment.ID = uuid.NewV4().String()
	comment.CreatedAt = time.Now()

	return &comment, nil
}
