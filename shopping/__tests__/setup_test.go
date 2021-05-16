package __tests__

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Couldn't load the .env file ;3")
	}
}

func TestSetup(t *testing.T) {
	db, err := ConnectDB()

	testComment(t, db)
}