package models

import (
	core "github.com/mazezen/gin-blog/db"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID         int    `gorm:"column:post_id;NOT NULL"`
	CommentContent string `gorm:"column:comment_content;NOT NULL"`
	Email          string `gorm:"column:email;NOT NULL"`
}

func (m *Comment) Comment() string {
	return "comment"
}

func init() {
	core.GetDB().AutoMigrate(&Comment{})
}
