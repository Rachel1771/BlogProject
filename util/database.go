package util

import (
	"Blog/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	driverName := "mysql"
	user := "root"
	password := "527113"
	host := "localhost"
	port := "3306"
	database := "blog"
	charset := "utf-8"
	loc := "Asia/Shanghai"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		user,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("无法打开数据库:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db

}
func GetDb() *gorm.DB {
	return DB
}
