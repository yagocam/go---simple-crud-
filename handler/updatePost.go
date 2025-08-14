package handler

import (
	"awesomeProject/schemas"
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
	if err := ctx.ShouldBindJSON(&request); err != nil {
		errors := GetValidationErrors(err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": errors})
		return
	}
	post := schemas.Post{}
	if err := db.First(&post, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{fmt.Sprintf("Id não encontrado %s", id): id})
		return
	}
	if err := db.Updates(&request).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{fmt.Sprintf("Error delete post %s", id): id})
		return
	}

	sendSuccess(ctx, "update-post", ToPostResponse(post))
}
