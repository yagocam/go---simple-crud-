package handler

import (
	"awesomeProject/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowPostHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id obrigatorio"})
		return
	}

	post := schemas.Post{}
	if err := db.First(&post, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{fmt.Sprintf("Id não encontrado %s", id): id})
		return
	}

	sendSuccess(ctx, "show-post", ToPostResponse(post))
}
