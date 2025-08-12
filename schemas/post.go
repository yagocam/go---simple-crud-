package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Content string   `json:"content"`
	Tags    []string `gorm:"type:text" json:"tags" gorm:"serializer:json"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Tags      []string  `gorm:"type:text" json:"tags" gorm:"serializer:json"`
}
