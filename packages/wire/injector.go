package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	DB      *gorm.DB
	Service *BData
}
