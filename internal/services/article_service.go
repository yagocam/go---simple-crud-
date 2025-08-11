package services

import (
	"awesomeProject/internal/models"

	"gorm.io/gorm"
)

type ArticleService struct {
	DB *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{DB: db}
}

func (s *ArticleService) Create(req models.CreateArticleRequest) (models.Article, error) {
	article := req.ToArticle()
	if err := s.DB.Create(&article).Error; err != nil {
		return models.Article{}, err
	}
	return article, nil
}

func (s *ArticleService) GetAll() ([]models.Article, error) {
	var articles []models.Article
	err := models.DB.Find(&articles).Error
	return articles, err
}

func (s *ArticleService) GetByID(id uint) (models.Article, error) {
	var article models.Article
	err := models.DB.First(&article, id).Error
	return article, err
}

func (s *ArticleService) Update(id uint, updated models.UpdateArticleRequest) (models.Article, error) {
	var article models.Article

	if err := s.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	if err := s.DB.Model(&article).Updates(updated).Error; err != nil {
		return article, err
	}

	if err := s.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (s *ArticleService) Delete(id uint) error {
	return models.DB.Delete(&models.Article{}, id).Error
}
