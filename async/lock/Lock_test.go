package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	var x int64
	var wg sync.WaitGroup
	var lock sync.Mutex
	start := time.Now()
	wg.Add(2)
	go add(&x, &wg, &lock)
	go add(&x, &wg, &lock)
	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(x)
}
