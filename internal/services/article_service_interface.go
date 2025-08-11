package services

import "awesomeProject/internal/models"

type ArticleServiceInterface interface {
	Create(article models.Article) (models.Article, error)
	GetAll() ([]models.Article, error)
	GetById(articleId uint) (models.Article, error)
	Update(articleId uint, article models.Article) (models.Article, error)
	Delete(articleId uint) error
}
