package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePostHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id obrigatorio"})
		return
	}
	request := UpdatePostRequest{}

	if err := db.First(&post, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{fmt.Sprintf("Id não encontrado %s", id): id})
		return
	}
	if err := db.Updates(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{fmt.Sprintf("Error delete post %s", id): id})
		return
	}

	sendSuccess(ctx, "delete-post", ToPostResponse(post))
}
