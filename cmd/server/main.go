package main

import (
	"Smradity/internal/config"
	"Smradity/internal/db"
	"Smradity/internal/logger"
	"Smradity/internal/router"
)

func main() {
	cfg := config.LoadConfig()
	logger.SetupLogger(cfg.LogLevel)
	db.SetupMgmConnection(cfg.DatabaseName, cfg.DatabaseURI)
	r := router.SetupRouters()

	err := r.Run(":" + cfg.ServerPort)
	if err != nil {
		return
	}
}
