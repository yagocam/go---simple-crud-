package handler

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreatePostRequest struct {
	Title   string   `json:"title" binding:"required,min=3"`
	Author  string   `json:"author" binding:"required"`
	Content string   `json:"content" binding:"required,min=10"`
	Tags    []string `json:"tags" binding:"required,min=1,dive,required,min=3"`
}

var validationMessages = map[string]string{
	"required": "é obrigatório",
	"min":      "não atende o tamanho mínimo",
	"max":      "excede o tamanho máximo",
	"email":    "não é um e-mail válido",
}

func TranslateError(fe validator.FieldError) string {
	field := fe.Field()
	tag := fe.Tag()
	param := fe.Param()

	switch tag {
	case "required":
		return fmt.Sprintf("%s é obrigatório", field)
	case "min":
		return fmt.Sprintf("%s deve ter pelo menos %s caracteres", field, param)
	case "max":
		return fmt.Sprintf("%s deve ter no máximo %s caracteres", field, param)
	case "email":
		return fmt.Sprintf("%s não é um e-mail válido", field)
	default:
		return fmt.Sprintf("%s falhou na validação '%s'", field, tag)
	}
}

func GetValidationErrors(err error) map[string]string {
	out := make(map[string]string)
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, fe := range ve {
			out[fe.Field()] = TranslateError(fe)
		}
	} else {
		out["error"] = err.Error()
	}

	return out
}
