package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Article struct {
	gorm.Model
	ArticleId string `gorm:"index:idx_article_id;unique;varchar(20);not null;"`
	Title     string `gorm:"index:idx_title;unique;varchar(150);not null;"`
	Author    string `gorm:"varchar(30);not null;"`
	Reference string `gorm:"varchar(300);"`
	Link      string `gorm:"varchar(300);"`
	View      uint   `gorm:"number;not null;default:0;"`
	Like      uint   `gorm:"number;not null;default:0;"`
	Content   string `gorm:"text;"`
	Tags      []Tag  `gorm:"many2many:article_tag;"`
}

type Tag struct {
	gorm.Model
	TagName   string    `gorm:"index:idx_tag_name;varchar(20);not null;"`
	ArticleId string    `gorm:"index:idx_article_id;varchar(20);"`
	Articles  []Article `gorm:"many2many:article_tag;"`
}

func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{DSN: "root:123456@tcp(127.0.0.1:3306)/maxblog?charset=utf8mb4&parseTime=True&loc=Local"}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	db.AutoMigrate(&Article{}, &Tag{})
}
