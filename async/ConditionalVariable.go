/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/7/30 18:30
 * @version     v1.0
 * @filename    ConditionalVariable.go
 * @description
 ***************************************************************************/
package main

//func main() {
//	cond := sync.NewCond(new(sync.Mutex))
//	condition := 0
//
//	// Consumer
//	go func() {
//		for {
//			cond.L.Lock()
//			for condition == 0 {
//				cond.Wait()
//			}
//			condition--
//			fmt.Printf("Consumer: %d\n", condition)
//			cond.Signal()
//			cond.L.Unlock()
//		}
//	}()
//
//	// Producer
//	for {
//		time.Sleep(time.Second)
//		cond.L.Lock()
//		for condition == 3 {
//			cond.Wait()
//		}
//		condition++
//		fmt.Printf("Producer: %d\n", condition)
//		cond.Signal()
//		cond.L.Unlock()
//	}
//}
