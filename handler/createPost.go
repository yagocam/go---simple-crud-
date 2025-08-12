package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "POST posts",
	})
}
