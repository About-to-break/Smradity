package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("auth")
	{
		users.GET("/")

	}
}
