package main

import (
	"Smradity/internal"
	"Smradity/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	router := internal.SetupRouters()

	err := router.Run(":" + cfg.ServerPort)
	if err != nil {
		return
	}
}
