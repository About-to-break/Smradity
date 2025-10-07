package config

import (
	"os"
	"strings"
)

type Config struct {
	ServerPort string
	LogLevel   string
}

func LoadConfig() *Config {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevel == "" {
		logLevel = "info"
	}
	return &Config{
		ServerPort: serverPort,
		LogLevel:   logLevel,
	}
}
