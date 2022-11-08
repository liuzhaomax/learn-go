package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for num := range increase(ctx) {
		fmt.Println(num)
		if num == 3 {
			break
		}
	}
}

func increase(ctx context.Context) <-chan int {
	ch := make(chan int)
	num := 1
	go func() {
		for {
			select {
			case ch <- num:
				num++
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}
