package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	SetupRoutes(r)
	return r
}