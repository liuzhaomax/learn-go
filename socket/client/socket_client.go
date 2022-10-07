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
