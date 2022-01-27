package service

import (
	"fmt"
	"gin-blog/constants"
	core "gin-blog/db"
	"gin-blog/models"
)

//CreateUserService
func CreateUserService(username, password, phone, email string, state int) bool {
	db := core.GetDB()

	user := models.User{
		Username: username,
		Password: password,
		Phone:    phone,
		Email:    email,
		State:    state,
	}

	tx := db.Create(&user)
	if tx.RowsAffected > 0 {
		return true
	}

	return false
}

//GetUsersService
func GetUsersService() (userList []*models.User) {
	core.GetDB().Find(&userList)
	return userList
}

//DeleteUserByIdService
func DeleteUserByIdService(id int) (err error, ok bool) {
	err = core.GetDB().First(&models.User{}, id).Error
	if err != nil {
		return err, false
	}

	err = core.GetDB().Delete(&models.User{}, id).Error
	return err, true
}

//GetUserByIdService
func GetUserByIdService(id int) (user []*models.User) {
	core.GetDB().First(&user, id)
	return
}

//UpdateUserByIdService
func UpdateUserByIdService(id int, username, password, phone,email string) (user models.User, ok bool, err error) {
	err = core.GetDB().First(&user, id).Error
	if err != nil {
		return user,false, err
	}

	user.Username = username
	user.Password = password
	user.Phone = phone
	user.Email = email

	err = core.GetDB().Save(&user).Error

	if err != nil {
		return user, false, err
	}

	return user,true, nil
}

//DisableUserByIdService
func DisableUserByIdService(id int) (user models.User, err error) {
	err = core.GetDB().First(&user, id).Error
	fmt.Println(user)
	if err != nil {
		return user, err
	}

	core.GetDB().Model(&user).Update("state", constants.DISABLE)

	return user, nil
}

//EnableUserByIdService
func EnableUserByIdService(id int) (user models.User, err error) {
	err = core.GetDB().First(&user, id).Error
	fmt.Println(user)
	if err != nil {
		return user, err
	}

	core.GetDB().Model(&user).Update("state", constants.ENABLE)

	return user, nil
}
