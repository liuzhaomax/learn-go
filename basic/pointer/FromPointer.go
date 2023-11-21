package pointer

import (
	"fmt"
)

// 有问题
func FromPointer(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	return value.(**interface{})
}

func Modify(s interface{}) {
	fmt.Println(**s.(**string))
}
