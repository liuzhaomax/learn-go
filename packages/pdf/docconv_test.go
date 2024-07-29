package pdf

import (
	"code.sajari.com/docconv/v2"
	"fmt"
	"testing"
)

func TestDocconv(t *testing.T) {
	res, err := docconv.ConvertPath("pdf17.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
