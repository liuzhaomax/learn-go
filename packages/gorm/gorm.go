package gorm

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

func InitDB() (*gorm.DB, func(), error) {
	logrus.Info("数据库连接启动")
	db, clean, err := LoadDB()
	if err != nil {
		logrus.Fatal("数据库连接失败", err)
		return nil, clean, err
	}
	err = AutoMigrate(db)
	if err != nil {
		logrus.Fatal("数据库表创建失败", err)
		return nil, clean, err
	}
	logrus.Info("数据库连接成功")
	createAdmin(db)
	return db, clean, err
}

func LoadDB() (*gorm.DB, func(), error) {
	gormLogger := logger.New(
		logrus.New(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			ParameterizedQueries:      true,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	logrus.Info(fmt.Sprintf("数据库品种: %s", "mysql"))
	db, err := gorm.Open(mysql.Open(DSN()), &gorm.Config{
		Logger: gormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	clean := func() {
		err = sqlDB.Close()
		if err != nil {
			logrus.Error("数据库断开连接失败", err)
		}
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, clean, err
	}
	sqlDB.SetMaxIdleConns(10000)
	sqlDB.SetMaxOpenConns(10000)
	sqlDB.SetConnMaxLifetime(time.Duration(3000) * time.Second)
	return db, clean, err
}

func AutoMigrate(db *gorm.DB) error {
	dbType := strings.ToLower("mysql")
	if dbType == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	err := db.AutoMigrate(new(Data))
	if err != nil {
		return err
	}
	return nil
}

func DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		"root", "123456", "127.0.0.1", "3306", "learn_go", "charset=utf8mb4&parseTime=True&loc=Local")
}

func createAdmin(db *gorm.DB) {
	var data Data
	result := db.First(&data)
	if result.RowsAffected == 0 {
		data.Mobile = "13012345678"
		data.Username = "admin"
		data.Password = "admin"
		res := db.Create(&data)
		if res.RowsAffected == 0 {
			fmt.Println(res.Error)
		}
	}
}

type Data struct {
	gorm.Model
	Mobile   string `gorm:"index:idx_mobile;unique;varchar(11);not null"`
	Username string `gorm:"index:idx_username;unique;varchar(50);not null"`
	Password string `gorm:"varchar(50);not null"`
}
