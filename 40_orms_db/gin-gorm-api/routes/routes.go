package routes

import (
	"gin-gorm-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api/v1")

	users := api.Group("/users")
	{
		users.POST("/", controllers.CreateUser)
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
	}
}
