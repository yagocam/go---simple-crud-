package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api/")
	{
		api.GET("/")
	}
}
