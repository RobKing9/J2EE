package dao

import (
	"Experiment1/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitMysql() (db2 *gorm.DB, err error) {
	dsn := "root:LKD2020.0921.@tcp/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	//自动创建数据表
	db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	return db, db.DB().Ping()
}

func GetDB() *gorm.DB {
	db, err := InitMysql()
	if err != nil {
		fmt.Println("get db failed:", err.Error())
	}
	return db
}
