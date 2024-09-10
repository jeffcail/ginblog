package service

import (
	core "github.com/jeffcail/gin-blog/db"
	"github.com/jeffcail/gin-blog/models"
)

// CreatePostService
func CreatePostService(title, desc, content, author, tags string, category int) (ok bool) {
	db := core.GetDB()

	posts := models.Posts{
		Title:    title,
		Desc:     desc,
		Content:  content,
		Author:   author,
		Tags:     tags,
		Category: category,
	}

	tx := db.Create(&posts)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetPostsListService
func GetPostsListService() (posts []*models.Posts) {
	db := core.GetDB()

	db.Find(&posts)

	return posts

}

// UpdatePostByIdService
func UpdatePostByIdService(id int, title, desc, content, author, tags string) (posts *models.Posts, flag int, err error) {
	db := core.GetDB()

	err = db.First(&posts, id).Error
	if err != nil {
		return posts, 0, nil
	}

	posts.Title = title
	posts.Desc = desc
	posts.Content = content
	posts.Author = author
	posts.Tags = tags

	tx := db.Model(&posts).Select("title", "desc", "content", "author", "tags").Updates(&posts)
	if tx.RowsAffected > 0 {
		return posts, 1, nil
	}
	return posts, 2, err
}

// DeletePostByIdService
func DeletePostByIdService(id int) (flag int, err error) {
	db := core.GetDB()

	err = db.First(&models.Posts{}, id).Error
	if err != nil {
		return 0, nil
	}

	err = db.Delete(&models.Posts{}, id).Error
	if err != nil {
		return 2, err
	}

	return 1, nil
}
