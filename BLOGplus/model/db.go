package model

import (
	"BLOGplus/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func InitDb() {
	var err error
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Println("连接数据库失败，请检查连接参数err:", err)
	}
	//禁用默认表名的复数形式
	db.SingularTable(true)
	//自动迁移表
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	//最大空闲
	db.DB().SetMaxIdleConns(100)
	//最大连接
	db.DB().SetMaxOpenConns(100)
	//最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)
	//db.Close()
	return
}
