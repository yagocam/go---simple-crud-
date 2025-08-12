package handler

import (
	"awesomeProject/schemas"

	"gorm.io/gorm"
)

type PostResponse struct {
	ID        uint           `json:"id"`
	CreatedAt MyTime         `json:"created_at"`
	UpdatedAt MyTime         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Content   string         `json:"content"`
	Tags      []string       `gorm:"type:text" json:"tags" gorm:"serializer:json"`
}

func ToPostResponse(post schemas.Post) PostResponse {
	return PostResponse{
		ID:        post.ID,
		CreatedAt: MyTime(post.CreatedAt),
		UpdatedAt: MyTime(post.UpdatedAt),
		DeletedAt: post.DeletedAt,
		Title:     post.Title,
		Author:    post.Author,
		Content:   post.Content,
		Tags:      post.Tags,
	}
}
