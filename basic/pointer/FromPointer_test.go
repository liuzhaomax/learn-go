package pointer

import (
	"fmt"
	"testing"
)

func TestFromPointer(t *testing.T) {
	var a = "1"
	b := &a
	c := FromPointer(&b)
	fmt.Println(c)
	//s := "ss"
	//a := &s
	//Modify(&a)
}
