package __tests__

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"))

	if err != nil {
		return nil, err
	}

	return db, nil
}
