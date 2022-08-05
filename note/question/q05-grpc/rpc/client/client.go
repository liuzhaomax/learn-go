/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/8/4 23:59
 * @version     v1.0
 * @filename    client.go
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	c, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	reply := ""

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))

	err = client.Call("FoodService.SayName", "锅包肉", &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)
}
