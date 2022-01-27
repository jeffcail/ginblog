package handler

import (
	"gin-blog/service"
	"gin-blog/types"
	"gin-blog/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreateTags
func CreateTags(c *gin.Context)  {
	name, ok := c.GetPostForm("name"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	_, ok, res := service.CreateTagsService(name); if !ok {
		util.Error(c, int(types.ApiCode.EXISTSNAME), types.ApiCode.GetMessage(types.ApiCode.EXISTSNAME))
		return
	}

	if !res {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}

//GetTagsList
func GetTagsList(c *gin.Context)  {
	tags := service.GetTagsListService()

	util.Success(c, tags)
}

//UpdateTagsById
func UpdateTagsById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	name, ok := c.GetPostForm("name"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	_, flag := service.UpdateTagsByIdService(newId, name)
	switch flag {
	case 0:
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
	case 2:
		util.Error(c, int(types.ApiCode.EXISTSNAME), types.ApiCode.GetMessage(types.ApiCode.EXISTSNAME))
	case 3:
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
	case 1:
		util.Success(c, nil)
	}
}

//DeleteTagsById
func DeleteTagsById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	_, flag := service.DeleteTagsByIdService(newId)

	switch flag {
	case 1:
		util.Success(c, nil)
	case 2:
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
	case 0:
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
	}
}
