package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/gin-blog/service"
	"github.com/jeffcail/gin-blog/types"
	"github.com/jeffcail/gin-blog/util"
	"strconv"
)

// CreatePost
func CreatePost(c *gin.Context) {
	title, ok := c.GetPostForm("title")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	desc := c.PostForm("desc")
	content, ok := c.GetPostForm("content")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	author, ok := c.GetPostForm("author")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	tags := c.PostForm("tags")
	category := c.PostForm("category")

	newCategory, err := strconv.Atoi(category)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	ok = service.CreatePostService(title, desc, content, author, tags, newCategory)
	if !ok {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}

// GetPostsList
func GetPostsList(c *gin.Context) {
	posts := service.GetPostsListService()

	util.Success(c, posts)
}

// UpdatePostById
func UpdatePostById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	title, ok := c.GetPostForm("title")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	desc := c.PostForm("desc")
	content, ok := c.GetPostForm("content")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	author, ok := c.GetPostForm("author")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	tags := c.PostForm("tags")

	_, flag, _ := service.UpdatePostByIdService(newId, title, desc, content, author, tags)
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

// DeletePostById
func DeletePostById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	flag, _ := service.DeletePostByIdService(newId)
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
