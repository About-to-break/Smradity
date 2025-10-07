package main

import (
	"Smradity/internal"
	"Smradity/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	internal.SetupLogger(cfg.LogLevel)
	router := internal.SetupRouters()

	err := router.Run(":" + cfg.ServerPort)
	if err != nil {
		return
	}
}
