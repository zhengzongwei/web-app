// Package          database
// @Title           mysql.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/15 17:03

package database

import (
	"backend/app/config"
	"backend/app/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var SQLDB *gorm.DB

func InitDB() {
	conf := config.Get()

	//fmt.Printf("%v", conf.MySQL.Read)
	mysqlConf := conf.MySQL.Read
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User,
		mysqlConf.Passwd,
		mysqlConf.Addr,
		mysqlConf.Port,
		mysqlConf.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败", err)
	}
	//db.SingularTable(true)
	////设置连接池
	////空闲
	//db.DB().SetMaxIdleConns(50)
	////打开
	//db.DB().SetMaxOpenConns(100)
	////超时
	//db.DB().SetConnMaxLifetime(time.Second * 30)

	SQLDB = db

	// 打印SQL语句
	//SQLDB = SQLDB.Debug()
	// 迁移数据库表
	err = SQLDB.AutoMigrate(
		&models.Book{},
		&models.Author{},
	)
	if err != nil {
		fmt.Println("迁移数据库失败", err)
	}
}

func GetDB() *gorm.DB {
	if SQLDB == nil {
		InitDB()
	}
	return SQLDB
}
