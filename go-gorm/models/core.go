package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:Kevin614718@tcp(127.0.0.1:3306)/poker_scorer?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = DB.Table("ps_user")
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println(err)
	}
}
