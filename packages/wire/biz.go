package main

import (
	"context"
	"github.com/google/wire"
	"learn-go/packages/wire/pb"
)

var BDataSet = wire.NewSet(wire.Struct(new(BData), "*"))

var ServiceSet = wire.NewSet(
	BDataSet,
)

type BData struct {
	MData *MData
}

func (bData *BData) GetDataById(ctx context.Context, req *pb.IdRequest) (*pb.DataRes, error) {
	var data *Data
	err := bData.MData.QueryDataById(req, data)
	if err != nil {
		return nil, err
	}
	res := Model2PB(data)
	return res, nil
}
