package handler

import (
	"awesomeProject/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPostHandler(ctx *gin.Context) {

	var posts []schemas.Post
	if err := db.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var response []PostResponse
	for _, p := range posts {
		response = append(response, ToPostResponse(p))
	}

	ctx.JSON(http.StatusOK, response)
}
