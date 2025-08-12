package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowPostHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "POST posts",
	})
}
