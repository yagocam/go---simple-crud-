package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(ctx *gin.Context) {
	request := CreatePostRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		errors := GetValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}
	post := request.ToPost()

	if err := db.Create(&post).Error; err != nil {
		logger.Errorf("create post error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	sendSuccess(ctx, "create-post", ToPostResponse(post))
}
