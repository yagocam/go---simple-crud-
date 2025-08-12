package schemas

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string         `json:"title"`
	Author  string         `json:"author"`
	Content string         `json:"content"`
	Tags    pq.StringArray `gorm:"type:text[]" json:"tags"`
}
