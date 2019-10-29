package routes

import (
	"github.com/tuananh1998hust/gin_tutorial/controllers"

	"github.com/gin-gonic/gin"
)

// SetUpRoutes :
func SetUpRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controllers.Ping)

	// API Routes
	v1 := r.Group("api/v1")
	{
		// Auth API
		v1.POST("/auth/register", controllers.CreateUser)
		v1.POST("auth/login", controllers.Login)
	}

	return r
}
