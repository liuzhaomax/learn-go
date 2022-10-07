package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstanceOfContext(t *testing.T) {
	context := GetInstanceOfContext()
	context.Work()
	context.Name = "123"

	ctx := new(Context)
	ctx.Work()

	fmt.Println(&ctx, &context)
	fmt.Println(GetInstanceOfContext().Name)
}
