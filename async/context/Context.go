package main

import (
	"context"
)

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
