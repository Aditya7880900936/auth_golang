package routes

import (
	"github.com/Aditya7880900936/auth_golang.git/controllers"
	"github.com/Aditya7880900936/auth_golang.git/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
