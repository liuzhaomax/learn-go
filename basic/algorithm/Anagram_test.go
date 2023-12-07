package algorithm

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"strconv"
	"testing"
)

func TestGetAnagram(t *testing.T) {
	s := "123456"
	fmt.Println(getAnagram(s))
	a := strconv.Itoa(int(codes.Unauthenticated))
	fmt.Println(a)
}
