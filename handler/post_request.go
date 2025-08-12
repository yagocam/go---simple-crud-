package handler

import "awesomeProject/schemas"

type CreatePostRequest struct {
	Title   string   `json:"title" binding:"required,min=3"`
	Author  string   `json:"author" binding:"required"`
	Content string   `json:"content" binding:"required,min=10"`
	Tags    []string `json:"tags" binding:"required,min=1,dive,required,min=3"`
}

type UpdatePostRequest struct {
	Title   *string   `json:"title" binding:"omitempty,min=3"`
	Author  *string   `json:"author" binding:"omitempty"`
	Content *string   `json:"content" binding:"omitempty,min=10"`
	Tags    *[]string `json:"tags" binding:"omitempty,min=1,dive,required,min=3"`
}

func (req *CreatePostRequest) ToPost() schemas.Post {
	return schemas.Post{
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
		Tags:    req.Tags,
	}
}
