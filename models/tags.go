package models

import (
	core "gin-blog-api/db"
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	ID   uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;NOT NULL"`
}

func (t *Tags) Tags() string {
	return "tags"
}

func init()  {
	core.GetDB().AutoMigrate(&Tags{})
}
