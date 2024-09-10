package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/gin-blog/service"
	"github.com/jeffcail/gin-blog/types"
	"github.com/jeffcail/gin-blog/util"
	"strconv"
)

// CreateComment
func CreateComment(c *gin.Context) {
	post_id, ok := c.GetPostForm("post_id")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	convert_post_id, err := strconv.Atoi(post_id)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	comment_content, ok := c.GetPostForm("comment_content")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	email, ok := c.GetPostForm("email")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	ok = service.CreateCommentService(convert_post_id, comment_content, email)
	if !ok {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}

// GetCommentList
func GetCommentList(c *gin.Context) {
	comments := service.GetCommentListService()

	util.Success(c, comments)
}

// DelCommentById
func DelCommentById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	convert_id, err := strconv.Atoi(id)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	flag, _ := service.DelCommentByIdService(convert_id)
	switch flag {
	case 0:
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	case 1:
		util.Success(c, nil)
	case 2:
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}
}
