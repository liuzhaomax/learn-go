package main

import (
	"fmt"
	"golang.org/x/net/context"
	"testing"
)

func TestIncrease(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for num := range increase(ctx) {
		fmt.Println(num)
		if num == 3 {
			break
		}
	}
}
