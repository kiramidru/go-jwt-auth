package routes

import (
	controller "jwt-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", controller.Signup())
	router.GET("/login", controller.Login())
}
