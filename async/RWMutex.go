package main

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	var lock sync.RWMutex
//	var wg sync.WaitGroup
//	var str = "hello"
//	wg.Add(4)
//	go readStr(&str, &wg, &lock, 0)
//	go readStr(&str, &wg, &lock, 0)
//	go readStr(&str, &wg, &lock, 1)
//	go updateStr(&str, &wg, &lock)
//	wg.Wait()
//	fmt.Println("over")
//}

func updateStr(str *string, wg *sync.WaitGroup, lock *sync.RWMutex) {
	lock.Lock()
	fmt.Println("写开始")
	*str = "world"
	fmt.Println("写结束")
	lock.Unlock()
	wg.Done()
}

func readStr(str *string, wg *sync.WaitGroup, lock *sync.RWMutex, sleep int) {
	lock.RLock()
	fmt.Println("读开始")
	time.Sleep(time.Second * time.Duration(sleep))
	fmt.Println(*str)
	fmt.Println("读结束")
	lock.RUnlock()
	wg.Done()
}

//写开始
//写结束
//读开始
//读开始
//world
//读结束
//读开始
//world
//读结束
//world
//读结束
//over
