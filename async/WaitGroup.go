package main

import (
	"fmt"
	"time"
)

//func main() {
//	var wg = sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		count(5, "sheep")
//		wg.Done()
//	}()
//	go func() {
//		count(3, "dog")
//		wg.Done()
//	}()
//	wg.Wait()
//}

func count(n int, animal string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Millisecond * 500)
	}
}
