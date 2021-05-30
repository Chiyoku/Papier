package model

type Comment struct {
	Identifer
	Message string `json:"message"`
	Author  string `json:"author"`
	Video   string `json:"video"`
}

func NewComment(message, author, video string) (*Comment, error) {
	comment := Comment{
		Message: message,
		Video: video,
		Author: author,
	}

	err := comment.Identifer.Create()

	if err != nil {
		return nil, err
	}

	return &comment, nil
}
