package router

import (
	"awesomeProject/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	api := router.Group("/api/")
	{
		api.GET("/post", handler.ShowPostHandler)
		api.GET("/posts", handler.ListPostHandler)
		api.POST("/post", handler.CreatePostHandler)
		api.DELETE("/post", handler.DeletePostHandler)
		api.PUT("/post", handler.UpdatePostHandler)
	}
}
