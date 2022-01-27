package handler

import (
	"fmt"
	"gin-blog/service"
	"gin-blog/types"
	"gin-blog/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

//CreateUser
func CreateUser(c *gin.Context) {
	username, ok := c.GetPostForm("username")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	password, ok := c.GetPostForm("password")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	newPassword := util.EncryMd5(password)

	phone := c.PostForm("phone")
	email := c.PostForm("email")
	state, _ := strconv.Atoi(c.PostForm("state"))

	ok = service.CreateUserService(username, newPassword, phone, email, state)
	if !ok {
		util.Error(c, int(types.ApiCode.CREATEUSERFAILED), types.ApiCode.GetMessage(types.ApiCode.CREATEUSERFAILED))
		return
	}
	util.Success(c, nil)

}
//FetchUsers
func GetUsers(c *gin.Context)  {
	userList := service.GetUsersService()
	fmt.Println(userList)

	util.Success(c, userList)
}

//DeleteUserById
func DeleteUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	fmt.Println("-------" + id + "---------")

	newId, err := strconv.Atoi(id)
	fmt.Printf("%T", newId)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	err, ok = service.DeleteUserByIdService(newId)
	if !ok {
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	if err != nil {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}

//GetUserById
func GetUserById(c *gin.Context)  {
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

	user := service.GetUserByIdService(newId)

	util.Success(c, user)
}

//UpdateUserById
func UpdateUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	username, ok := c.GetPostForm("username"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	password, ok := c.GetPostForm("password"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newPassword := util.EncryMd5(password)

	phone := c.PostForm("phone")
	email := c.PostForm("email")

	_, ok, err = service.UpdateUserByIdService(newId, username, newPassword, phone, email); if !ok {
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	if err != nil {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}

//DisableUserById
func DisableUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	_, err = service.DisableUserByIdService(newId); if err != nil {
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	util.Success(c, nil)
}

//EnableUserById
func EnableUserById(c *gin.Context)  {
	id, ok := c.GetPostForm("id"); if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id); if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	_, err = service.EnableUserByIdService(newId); if err != nil {
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	util.Success(c, nil)
}
