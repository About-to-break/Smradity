package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	ServerPort   string
	LogLevel     string
	DatabaseName string
	DatabaseURI  string
}

func LoadConfig() *Config {
	slog.Info("Loading config...")

	if err := godotenv.Load(".env"); err != nil {
		slog.Warn(".env file was not used in configuration")
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevel == "" {
		logLevel = "info"
	}
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		slog.Error("No DATABASE_NAME specified")
	}
	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI == "" {
		slog.Error("No DATABASE_URI specified")
	}
	slog.Info("Loading config successful")
	return &Config{
		ServerPort:   serverPort,
		LogLevel:     logLevel,
		DatabaseName: databaseName,
		DatabaseURI:  databaseURI,
	}
}
