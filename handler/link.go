package handler

import (
	"gin-blog-api/service"
	"gin-blog-api/types"
	"gin-blog-api/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreateLink
func CreateLink(c *gin.Context)  {
	name, ok := c.GetPostForm("name"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	url, ok := c.GetPostForm("url"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	ok = service.CreateLinkService(name, url); if !ok {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}

//GetLinkList
func GetLinkList(c *gin.Context)  {
	links := service.GetLinkListService()

	util.Success(c, links)
}

//UpdateLinkById
func UpdateLinkById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	convert_id, err := strconv.Atoi(id); if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	name, ok := c.GetPostForm("name"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	url, ok := c.GetPostForm("url"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	_, flag := service.UpdateLinkByIdService(convert_id, name, url)
	switch flag {
	case 0:
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	case 1:
		util.Success(c, nil)
		return
	case 2:
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}
}

//DeleteLinkById
func DeleteLinkById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	convert_id, err := strconv.Atoi(id); if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	res := service.DeleteLinkByIdService(convert_id)
	switch res {
	case 0:
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	case 1:
		util.Success(c, nil)
		return
	case 2:
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}
}
