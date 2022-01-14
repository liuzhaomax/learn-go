/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/6 20:54
 * @version     v1.0
 * @filename    Channel.go
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
	"time"
)

//func main() {
//	c := make(chan string)
//	go count1(5, "sheep", c)
//	for message := range c {
//		fmt.Println(message)
//	}
//}

func count1(n int, animal string, c chan string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
