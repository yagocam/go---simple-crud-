package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api/")
	{
		api.GET("/post", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "get post",
			})
		})
		api.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "get posts",
			})
		})
		api.POST("/post", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "POST posts",
			})
		})
		api.PUT("/post", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "PUT posts",
			})
		})
		api.DELETE("/post", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "DEL posts",
			})
		})
	}
}
