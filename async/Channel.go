package main

import (
	"fmt"
	"time"
)

//func main() {
//	c := make(chan string)
//	go count1(5, "sheep", c)
//	for message := range c {
//		fmt.Println(message)
//	}
//}

func count1(n int, animal string, c chan string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
