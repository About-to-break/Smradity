package internal

import (
	"github.com/gin-gonic/gin"

	"Smradity/internal/auth"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	appRoutes := router.Group("/")
	{
		auth.AuthRoutes(appRoutes)
	}

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "Hello World"})
	})

	return router
}
