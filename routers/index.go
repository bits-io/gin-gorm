package routers

import (
	"github.com/gin-gonic/gin"
	"gin_gorm/controllers"
	"gin_gorm/routers/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())

	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/books", controllers.GetBooks)
	}
}