package main

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
			go WorkCat(chDog, "cat")
		case <-chDog:
			go WorkDog(chFish, "dog")
		case <-chFish:
			go WorkFish(chCat, "fish")
		}
	}
}

func WorkCat(ch chan bool, animal string) {
	fmt.Println(animal)
	ch <- true
}

func WorkDog(ch chan bool, animal string) {
	fmt.Println(animal)
	ch <- true
}

func WorkFish(ch chan bool, animal string) {
	fmt.Println(animal)
	time.Sleep(time.Second)
	ch <- true
}
