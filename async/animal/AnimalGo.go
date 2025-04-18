package animal

import (
	"context"
	"fmt"
	"time"
)

func AnimalGo(ctx context.Context, times int) {
	chCat := make(chan bool, 1)
	chDog := make(chan bool, 1)
	chFish := make(chan bool, 1)

	chCat <- true // 启动信号

	for i := 0; i < times; i++ {
		select {
		case <-ctx.Done():
			return
		case <-chCat:
			Animal(chDog, "cat")
		case <-chDog:
			Animal(chFish, "dog")
		case <-chFish:
			Animal(chCat, "fish")
		}
	}
}

func Animal(ch chan bool, animal string) {
	fmt.Println(animal)
	if animal == "fish" {
		time.Sleep(time.Second)
	}
	ch <- true
}
