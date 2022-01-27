package models

import (
	core "gin-blog-api/db"
	"gorm.io/gorm"
)

//User
type User struct {
	gorm.Model
	Username string    `gorm:"column:username;NOT NULL"`
	Password string    `gorm:"column:password;NOT NULL"`
	Phone    string    `gorm:"column:phone"`
	Email    string    `gorm:"column:email"`
	State    int       `gorm:"column:state;default:1;NOT NULL"`
}

//User
func (u *User) User() string {
	return "user"
}

func init()  {
	core.Db.AutoMigrate(&User{})
}
