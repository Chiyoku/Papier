package tests

import (
	"log"
	"shopping/internal/factory"
	"shopping/internal/model"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Couldn't connect to the database cause >>> %s\n", err)
	}

	db.AutoMigrate(&model.Comment{})

	return db
}

func TestSetup(t *testing.T) {
	var (
		db = connectDB()
		commentRepo = factory.CommentRepoFactory(db)
		commentUseCase = factory.CommentUseCaseFactory(&commentRepo)
	)

	testCommentRepo(t, commentRepo)
	testCommentUseCase(t, commentUseCase)
}
