package algorithm

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"testing"
)

type CODE uint32

const ABC CODE = 10

func (c CODE) String() string {
	switch c {
	case ABC:
		return "ABC"
	default:
		return "CODE"
	}
}

func TestName(t *testing.T) {
	fmt.Println(codes.Aborted) // const Aborted Code = 10
	fmt.Println(int(codes.Aborted))
	fmt.Println(ABC)
}
