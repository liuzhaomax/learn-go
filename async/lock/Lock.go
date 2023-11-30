package lock

import (
	"sync"
)

func add(x *int64, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 如果不加锁，只有很小很小的概率会加到10000
		*x++
		lock.Unlock()
	}
	wg.Done()
}
