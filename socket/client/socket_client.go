/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/13 2:28
 * @version     v1.0
 * @filename    socket_client.go
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("net error", err)
		return
	}
	fmt.Println("connection established")

	for {
		sendData := []byte("hello world")
		cnt, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("net error", err)
			return
		}
		fmt.Println("client ===> server cnt:", cnt, ", data:", string(sendData[:cnt]))

		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("net error", err)
			return
		}
		fmt.Println("client <===server cnt:", cnt, ", data:", string(buf[:cnt]))
		time.Sleep(time.Second)
	}
	//_ = conn.Close()
}
