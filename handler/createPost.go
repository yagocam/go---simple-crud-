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
	}
	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("create post error: %v", err)
		return
	}
}
