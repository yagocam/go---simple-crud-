package handler

import (
	"awesomeProject/schemas"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListPostHandler(ctx *gin.Context) {

	var posts []schemas.Post
	tagsParam := ctx.Query("tags")
	var tags []string
	if tagsParam != "" {
		tags = strings.Split(tagsParam, ",")
	}

	query := db.Model(&schemas.Post{})

	if len(tags) > 0 {
		for _, tag := range tags {
			query = query.Where("tags LIKE ?", "%"+tag+"%")
		}
	}
	if err := query.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var response []PostResponse
	for _, p := range posts {
		response = append(response, ToPostResponse(p))
	}

	ctx.JSON(http.StatusOK, response)
}
