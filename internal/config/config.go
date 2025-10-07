package config

import (
	"os"
)

type Config struct {
	ServerPort string
}

func LoadConfig() *Config {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	return &Config{
		ServerPort: serverPort,
	}
}
