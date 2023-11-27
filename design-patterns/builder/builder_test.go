package builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	concrete := NewConcreteBuilder("con").setBuilt(true).Build()
	result := concrete.GetResult() // concrete不具备set方法
	fmt.Println(result.Name)
	fmt.Println(result.Built)
}
