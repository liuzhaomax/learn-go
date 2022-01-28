/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/8 4:54
 * @version     v1.0
 * @filename    Lock.go
 * @description
 ***************************************************************************/
package main

import (
	"sync"
)

var x int64

//var wg sync.WaitGroup
var lock sync.Mutex

//func main() {
//	start := time.Now()
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(time.Since(start))
//	fmt.Println(x)
//}

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}
