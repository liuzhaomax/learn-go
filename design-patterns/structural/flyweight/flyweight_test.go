package flyweight

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFlyWeightFactory_GetFlyWeight(t *testing.T) {
	factory := NewFlyWeightFactory()
	hong := factory.GetFlyWeight("hong hong hong", 10)
	lv := factory.GetFlyWeight("lv lv lv", 20)
	lv2 := factory.GetFlyWeight("lv lv lv", 30)

	assert.Equal(t, 2, len(factory.pool))
	assert.Equal(t, &hong, factory.pool["hong hong hong"])
	assert.Equal(t, &lv, factory.pool["lv lv lv"]) // 共享部分被修改
	assert.Equal(t, &lv2, factory.pool["lv lv lv"])

	fmt.Println(lv)
	fmt.Println(lv2)
	fmt.Println(factory)
}
