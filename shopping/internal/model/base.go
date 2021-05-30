package model

import (
	"time"

	"github.com/myesui/uuid"
)

type Identifer struct {
	ID string `json:"id" gorm:"type:uuid;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (identifer *Identifer) Create() error {
	identifer.ID = uuid.NewV4().String()
	identifer.CreatedAt = time.Now()
	identifer.UpdatedAt = time.Now()

	return nil
}
