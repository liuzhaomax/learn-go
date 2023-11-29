package responsibility_chain

import (
	"fmt"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	wang := NewHandler("wang", nil, 1)
	zhang := NewHandler("zhang", wang, 2)
	liu := NewHandler("liu", zhang, 3)

	r := wang.Handle(1)
	fmt.Println(r)
	r = zhang.Handle(2)
	fmt.Println(r)
	r = liu.Handle(2)
	fmt.Println(r)
}
