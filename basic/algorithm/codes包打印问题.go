package algorithm

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(codes.Aborted) // const Aborted Code = 10
	fmt.Println(int(codes.Aborted))
	fmt.Println(codes.Aborted == 10)
}
