package main

import (
	"errors"
	"github.com/google/wire"
	"gorm.io/gorm"
	"learn-go/packages/wire/pb"
)

var MDataSet = wire.NewSet(wire.Struct(new(MData), "*"))

var ModelSet = wire.NewSet(
	MDataSet,
)

type Data struct {
	gorm.Model
	Mobile string `gorm:"index:idx_mobile;unique;varchar(11);not null"`
}

func Model2PB(data *Data) *pb.DataRes {
	dataRes := &pb.DataRes{
		Id:     int32(data.ID),
		Mobile: data.Mobile,
	}
	return dataRes
}

type MData struct {
	DB *gorm.DB
}

func (mData *MData) QueryDataById(req *pb.IdRequest, data *Data) error {
	result := mData.DB.First(&data, req.Id)
	if result.RowsAffected == 0 {
		return errors.New("数据没找到")
	}
	return nil
}
