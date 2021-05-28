package server

import (
	"fmt"
	"os"
)

type ServerConfig struct {
	Address   string
	Port      string
	SecretKey string
}

func CreateDefaultConfig() *ServerConfig {
	address := os.Getenv("ADDR")
	port := os.Getenv("PORT")
	secretKey := os.Getenv("SECRET_KEY")

	if port == "" {
		port = "4040"
	}
	if address == "" {
		address = "localhost"
	}
	return &ServerConfig{
		Address:   address,
		Port:      port,
		SecretKey: secretKey,
	}
}

func (config *ServerConfig) AddrToString() string {
	return fmt.Sprintf("%s:%s", config.Address, config.Port)
}
