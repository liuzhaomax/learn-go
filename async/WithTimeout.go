package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context.WithTimeout

func worker(ctx context.Context, wg *sync.WaitGroup) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Second * 1) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

//func main() {
//	var wg sync.WaitGroup
//	// 设置一个50毫秒的超时
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
//	wg.Add(1)
//	go worker(ctx, &wg)
//	time.Sleep(time.Second * 6)
//	cancel() // 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")
//}
