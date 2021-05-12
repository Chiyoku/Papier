package server

import (
	"fmt"
	"os"
)

type ServerConfig struct {
	Address string
	Port    string
}

func CreateDefaultConfig() *ServerConfig {
	address := os.Getenv("ADDR")
	port := os.Getenv("PORT")
	if port == "" {
		port = "4040"
	}
	if address == "" {
		address = "localhost"
	}
	return &ServerConfig{
		Address: address,
		Port:    port,
	}
}

func (config *ServerConfig) ToString() string {
	return fmt.Sprintf("%s:%s", config.Address, config.Port)
}
