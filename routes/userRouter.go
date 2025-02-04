package routes

import (
	controller "jwt-auth/controllers"
	"jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.GET("/user/:user_id", controller.GetUser())
	router.GET("/users", controller.GetUsers())
}
