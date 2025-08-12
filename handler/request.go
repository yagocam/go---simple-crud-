package handler

type CreatePostRequest struct {
	Title   string   `json:"title" binding:"required,min=3"`
	Author  string   `json:"author" binding:"required"`
	Content string   `json:"content" binding:"required,min=10"`
	Tags    []string `json:"tags" binding:"required,min=1,dive,required,min=3"`
}
