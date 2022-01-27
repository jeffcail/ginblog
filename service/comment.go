package service

import (
	core "gin-blog/db"
	"gin-blog/models"
)

// CreateCommentService
func CreateCommentService(post_id int, comment_content string, email string) bool {
	db := core.GetDB()

	tx := db.Create(&models.Comment{
		PostID: post_id,
		CommentContent: comment_content,
		Email: email,
	})

	if tx.RowsAffected > 0 {
		return true
	}

	return false
}

// GetCommentListService
func GetCommentListService() (comments []*models.Comment) {
	db := core.GetDB()

	db.Find(&comments)

	return comments
}

// DelCommentByIdService
func DelCommentByIdService(id int) (flag int, err error) {
	db := core.GetDB()

	comment := models.Comment{}

	err = db.First(&comment, id).Error; if err != nil {
		return 0, nil
	}

	err = db.Delete(&comment, id).Error; if err != nil {
		return 2, err
	}

	return 1, nil
}
