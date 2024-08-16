package main

import (
	"fmt"
	"reflect"
	"testing"
)

type UserWechatInfo struct {
	ID         string `gorm:"primarykey;size:22"`
	Openid     string `gorm:"index:idx_openid;not null;unique;size:50"`
	Nickname   string `gorm:"size:50;not null"`
	Sex        int    `gorm:"not null;default:0"`
	Province   string `gorm:"size:255"`
	City       string `gorm:"size:255"`
	Country    string `gorm:"size:255"`
	Headimgurl string `gorm:"size:255"`
	Privilege  string `gorm:"type:json"`
	UnionID    string `gorm:"size:50"`
}

func CopyStructFields(dest, src any) {
	srcVal := reflect.ValueOf(src).Elem()
	destVal := reflect.ValueOf(dest).Elem()
	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Type().Field(i)
		destField := destVal.FieldByName(field.Name)
		if destField.CanSet() {
			destField.Set(srcVal.Field(i))
		}
	}
}

func TestCopyStruct(t *testing.T) {
	userWechatInfo := &UserWechatInfo{
		ID:         "123",
		Openid:     "",
		Nickname:   "",
		Sex:        0,
		Province:   "",
		City:       "",
		Country:    "",
		Headimgurl: "",
		Privilege:  "",
		UnionID:    "",
	}
	mapUserWechatInfo := &UserWechatInfo{
		ID:         "456",
		Openid:     "",
		Nickname:   "",
		Sex:        1,
		Province:   "",
		City:       "",
		Country:    "",
		Headimgurl: "",
		Privilege:  "",
		UnionID:    "111",
	}
	CopyStructFields(userWechatInfo, mapUserWechatInfo)
	fmt.Println(userWechatInfo.ID, userWechatInfo.Sex, userWechatInfo.UnionID)
}
