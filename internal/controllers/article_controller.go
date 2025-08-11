package controllers

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	service services.ArticleService
}

func NewArticleController(service services.ArticleService) *ArticleController {
	return &ArticleController{service: service}
}

// Create - cria um novo artigo
func (c *ArticleController) Create(ctx *gin.Context) {
	var req models.CreateArticleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := c.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, article)
}

// List - lista todos os artigos
func (c *ArticleController) List(ctx *gin.Context) {
	articles, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

// GetByID - busca um artigo pelo ID
func (c *ArticleController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	article, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Artigo não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, article)
}

// Update - atualiza parcialmente um artigo
func (c *ArticleController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req models.UpdateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := c.service.Update(uint(id), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, article)
}

// Delete - remove um artigo pelo ID
func (c *ArticleController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
