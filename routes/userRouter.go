package routes

import "github.com/gin-gonic/gin"

// controller "go-jwt/controllers"
// "go-jwt/middleware"
// "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users".controller.GetUsers())
	incomingRoutes.GET("/users/:user_id".controller.GetUser())
}
