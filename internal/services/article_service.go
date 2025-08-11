package services

import (
	"awesomeProject/internal/models"

	"gorm.io/gorm"
)

type ArticleService interface {
	Create(article models.CreateArticleRequest) (models.Article, error)
	GetAll() ([]models.Article, error)
	GetById(articleId uint) (models.Article, error)
	Update(articleId uint, article models.UpdateArticleRequest) (models.Article, error)
	Delete(articleId uint) error
}
type articleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{db: db}
}

func (s *articleService) Create(req models.CreateArticleRequest) (models.Article, error) {
	article := req.ToArticle()
	if err := s.db.Create(&article).Error; err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func (s *articleService) GetAll() ([]models.Article, error) {
	var articles []models.Article
	err := models.DB.Find(&articles).Error
	return articles, err
}

func (s *articleService) GetById(id uint) (models.Article, error) {
	var article models.Article
	err := models.DB.First(&article, id).Error
	return article, err
}

func (s *articleService) Update(id uint, updated models.UpdateArticleRequest) (models.Article, error) {
	var article models.Article

	if err := s.db.First(&article, id).Error; err != nil {
		return article, err
	}

	if err := s.db.Model(&article).Updates(updated).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (s *articleService) Delete(id uint) error {
	return models.DB.Delete(&models.Article{}, id).Error
}
