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

//func main() {
//	var x int64
//	var wg sync.WaitGroup
//	var lock sync.Mutex
//	start := time.Now()
//	wg.Add(2)
//	go add(&x, &wg, &lock)
//	go add(&x, &wg, &lock)
//	wg.Wait()
//	fmt.Println(time.Since(start))
//	fmt.Println(x)
//}

func add(x *int64, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 如果不加锁，只有很小很小的概率会加到10000
		*x++
		lock.Unlock()
	}
	wg.Done()
}
