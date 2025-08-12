package schemas

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Content string   `json:"content"`
	Tags    []string `gorm:"type:text" json:"tags" gorm:"serializer:json"`
}
