/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/29 17:55
 * @version     v1.0
 * @filename    Fib.go
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
	"math"
	"time"
)

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	var skip int
	var total int
	_, err := fmt.Scanln(&skip)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&total)
	if err != nil {
		return
	}
	go Fib(resultChan, skip, total)
LOOP:
	for {
		select {
		case res := <-resultChan:
			if res == 0 {
				break LOOP
			}
			fmt.Println(res)
		}
	}
}

func Fib(resultChan chan int, skip int, total int) {
	defer close(resultChan)
	x, y := 1, 1
	count := 0
	for {
		x, y = y, x+y
		if count >= skip {
			resultChan <- x % int(math.Pow(10, 9))
			time.Sleep(time.Millisecond * 10)
		}
		if count+1 == skip+total {
			break
		}
		count++
	}
}

func main() {
	requestChan := make(chan bool)
	resultChan := make(chan int)
	ModuloFibonacciSequence(requestChan, resultChan)

}
