package handler

import (
	"awesomeProject/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePostHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	post := schemas.Post{}
	if err := db.First(&post, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{fmt.Sprintf("Id não encontrado %s", id): id})
		return
	}
	if err := db.Delete(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{fmt.Sprintf("Error delete post %s", id): id})
		return
	}

	sendSuccess(ctx, "delete-post", ToPostResponse(post))
}
