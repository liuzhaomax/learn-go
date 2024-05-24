package practice

import (
	"fmt"
	"sync"
)

func animal(wg *sync.WaitGroup, ch chan string) {
	var value string
	defer wg.Done()
	go func() {
		for {
			select {
			case value = <-ch:
				fmt.Println(value)
				switch value {
				case "cat":
					ch <- "dog"
					return
				case "dog":
					ch <- "fish"
					return
				default:
					ch <- "cat"
					return
				}
			default:
				fmt.Println("wrong")
				return
			}
		}
	}()
}

func run() {
	var wg sync.WaitGroup
	ch := make(chan string, 1)
	// defer close(ch)
	animals := []string{"cat", "dog", "fish"}
	ch <- "cat"
	for i := 0; i < 5; i++ {
		wg.Add(3)
		for range animals {
			animal(&wg, ch)
		}
		wg.Wait()
	}
}
