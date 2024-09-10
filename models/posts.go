package models

import (
	core "github.com/jeffcail/gin-blog/db"
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title    string `gorm:"column:title;NOT NULL"`
	Desc     string `gorm:"column:desc"`
	Content  string `gorm:"column:content;NOT NULL"`
	Author   string `gorm:"column:author"`
	Tags     string `gorm:"column:tags"`
	Category int    `gorm:"column:category"`
	Love     int    `gorm:"column:love"`
	Watch    int    `gorm:"column:watch"`
}

func (p *Posts) Posts() string {
	return "posts"
}

func init() {
	core.GetDB().AutoMigrate(&Posts{})
}
