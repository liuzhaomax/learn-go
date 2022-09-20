package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() (*gorm.DB, func(), error) {
	conn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "localhost", 3306, "maxblog-be-template")
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	clean := func() {
		fmt.Println("关闭")
	}
	return db, clean, nil
}
