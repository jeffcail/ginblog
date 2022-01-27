package service

import (
	core "gin-blog/db"
	"gin-blog/models"
)

//CreateCateService
func CreateCateService(name string) (flag int, err error) {
	db := core.GetDB()

	cate := models.Category{}

	err = db.Where("name = ?", name).First(&cate).Error
	if err == nil {
		return 0, err
	}

	cate.Name = name

	err = db.Create(&cate).Error
	if err != nil {
		return 2, err
	}

	return 1, nil
}

//GetCateListService
func GetCateListService() (cate []*models.Category) {
	db := core.GetDB()

	db.Find((&cate))

	return cate
}

//UpdateCateByIdService
func UpdateCateByIdService(id int, name string) (flag int, err error) {
	db := core.GetDB()

	cate := models.Category{}

	err = db.First(&cate, id).Error
	if err != nil {
		return 0, err
	}

	cate.Name = name

	err = db.Model(&cate).Update("name", name).Error
	if err != nil {
		return 2, err
	}

	return 1, err
}

//DeleteCateByIdService
func DeleteCateByIdService(id int) (flag int, err error) {
	db := core.GetDB()

	cate := models.Category{}

	err = db.First(&cate, id).Error; if err != nil {
		return 0, nil
	}

	err = db.Delete(&cate, id).Error; if err != nil {
		return 2, err
	}

	return 1, nil
}
