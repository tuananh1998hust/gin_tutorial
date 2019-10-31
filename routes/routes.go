package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tuananh1998hust/gin_tutorial/controllers"
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
		v1.GET("/user", controllers.GetUser)
		v1.PUT("/user/:id", controllers.UpdateUser)
	}

	return r
}
