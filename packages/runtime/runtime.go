package runtime

import (
	"fmt"
	"reflect"
)

func Run() {
	var a = new([]int)
	fmt.Println(reflect.TypeOf(a))
	b := append(*a, 0)
	fmt.Println(b[0])
}
