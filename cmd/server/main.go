package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "Hello World"})
	})

	err := router.Run(":9080")
	if err != nil {
		return
	}
}
