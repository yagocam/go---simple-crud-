package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Content string   `json:"content"`
	Tags    []string `json:"tags" gorm:"type:text[]"`
}

type CreateArticleRequest struct {
	Title   string   `json:"title" binding:"required,min=3"`
	Author  string   `json:"author" binding:"required"`
	Content string   `json:"content" binding:"required,min=10"`
	Tags    []string `json:"tags" binding:"required,min=1,dive,required,min:3"`
}
type UpdateArticleRequest struct {
	Title   *string   `json:"title" binding:"omitempty,min=3"`
	Author  *string   `json:"author" binding:"omitempty"`
	Content *string   `json:"content" binding:"omitempty,min=10"`
	Tags    *[]string `json:"tags" binding:"omitempty,min=1,dive,required,min=3"`
}

func (req *CreateArticleRequest) ToArticle() Article {
	return Article{
		Title:   req.Title,
		Author:  req.Author,
		Content: req.Content,
		Tags:    req.Tags,
	}
}
