package main

import (
	"Smradity/internal"
	"Smradity/internal/config"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()
	router := internal.SetupRouters()

	port := fmt.Sprintf(":%d", cfg.ServerPort)

	err := router.Run(port)
	if err != nil {
		return
	}
}
