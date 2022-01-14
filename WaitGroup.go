/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/6 20:44
 * @version     v1.0
 * @filename    WaitGroup.go
 * @description
 ***************************************************************************/
package main

//func main() {
//	var wg = sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		count(5, "sheep")
//		wg.Done()
//	}()
//	go func() {
//		go count(3, "dog")
//		wg.Done()
//	}()
//	wg.Wait()
//}
//
//func count(n int, animal string) {
//	for i := 0; i < n; i++ {
//		fmt.Println(i+1, animal)
//		time.Sleep(time.Millisecond * 500)
//	}
//}
