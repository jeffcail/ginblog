package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 定义db全局变量
var Db *gorm.DB

func init() {
	var err error
	dsn := "root:asdasdzxc@tcp(127.0.0.1:3306)/gin-blog-api?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		fmt.Println(err)
		panic("Connecting database failed: " + err.Error())
	}
}

// GetDB
func GetDB() *gorm.DB {
	return Db
}
