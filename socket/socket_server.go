package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "127.0.0.1"
	port := "8848"
	address := fmt.Sprintf("%s:%s", ip, port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net error", err)
		return
	}

	for {
		fmt.Println("listening...")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net error", err)
			return
		}
		fmt.Println("connection established! ")

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		fmt.Println("ready to receive")

		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("net error", err)
			return
		}
		fmt.Println("client ===> server, long:", cnt, ", data:", string(buf[:cnt]))

		cnt, err = conn.Write(buf)
		if err != nil {
			fmt.Println("net error", err)
			return
		}
		fmt.Println("client <=== server, long:", cnt, ", data:", string(buf[:cnt]))

		//_ = conn.Close()
	}
}
