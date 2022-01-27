package models

import (
	core "gin-blog/db"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Name string `gorm:"column:name;NOT NULL"`
	Url  string `gorm:"column:url;NOT NULL"`
}

func (m *Link) Link() string {
	return "link"
}

func init()  {
	db := core.GetDB()

	db.AutoMigrate(&Category{})
}
