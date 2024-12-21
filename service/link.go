package service

import (
	core "github.com/mazezen/gin-blog/db"
	"github.com/mazezen/gin-blog/models"
)

// CreateLinkService
func CreateLinkService(name, url string) bool {
	db := core.GetDB()

	tx := db.Create(&models.Link{Name: name, Url: url})

	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetLinkListService
func GetLinkListService() (links []*models.Link) {
	db := core.GetDB()

	db.Find(&links)

	return
}

// UpdateLinkByIdService
func UpdateLinkByIdService(id int, name, url string) (link *models.Link, flag int) {
	db := core.GetDB()

	err := db.First(&link, id).Error
	if err != nil {
		return link, 0
	}

	link.Name = name
	link.Url = url

	tx := db.Model(&link).Select("name", "url").Updates(link)
	if tx.RowsAffected > 0 {
		return link, 1
	}

	return link, 2
}

// DeleteLinkByIdService
func DeleteLinkByIdService(id int) int {
	db := core.GetDB()

	err := db.First(&models.Link{}, id).Error
	if err != nil {
		return 0
	}

	tx := db.Delete(&models.Link{}, id)
	if tx.RowsAffected > 0 {
		return 1
	}
	return 2
}
