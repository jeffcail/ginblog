package service

import (
	"fmt"
	core "gin-blog-api/db"
	"gin-blog-api/models"
)

//CreateTagsService
func CreateTagsService(name string) (tags *models.Tags, ok bool, res bool) {
	db := core.GetDB()

	err := db.Where("name = ?", name).First(&tags).Error
	if err == nil {
		return tags, false, false
	}
	tags.Name = name

	tx := db.Create(&tags)
	fmt.Println(tx.RowsAffected)
	if tx.RowsAffected > 0 {
		return tags, true, true
	}

	return tags, false, false
}

//GetTagsListService
func GetTagsListService() (tags []*models.Tags){
	db := core.GetDB()

	db.Find((&tags))
	return tags
}

//UpdateTagsByIdService
func UpdateTagsByIdService(id int, name string) (tags *models.Tags, flag int) {
	db := core.GetDB()

	err := db.First(&tags, id).Error; if err != nil {
		flag = 0
		return tags, flag
	}

	err = db.Where("name = ?", name).First(&tags).Error; if err == nil {
		flag = 2
		return tags, flag
	}

	tags.Name = name

	err = core.GetDB().Model(&tags).Save(&tags).Error; if err != nil {
		flag = 3
		return tags, flag
	}
	flag = 1
	return tags, flag
}

//DeleteTagsByIdService
func DeleteTagsByIdService(id int) (tags *models.Tags, flag int) {
	db := core.GetDB()

	err := db.First(&tags, id).Error; if err != nil {
		return tags, 0
	}

	err = db.Delete(&tags, id).Error; if err != nil {
		return tags, 2
	}

	return tags, 1
}
