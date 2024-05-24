package animal

import (
	"fmt"
	"time"
)

func AnimalGo() {
	chCat := make(chan bool, 1)
	chDog := make(chan bool, 1)
	chFish := make(chan bool, 1)

	chCat <- true

	for {
		select {
		case <-chCat:
			go Animal(chDog, "cat")
		case <-chDog:
			go Animal(chFish, "dog")
		case <-chFish:
			go Animal(chCat, "fish")
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
