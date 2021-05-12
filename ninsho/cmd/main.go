package main

import (
	"ninsho/internal/models"
	"ninsho/internal/server"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	config := server.CreateDefaultConfig()

	server_conn := server.NewServer(config, db)
	server_conn.Serve()
}
