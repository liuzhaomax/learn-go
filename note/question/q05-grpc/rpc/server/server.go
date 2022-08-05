/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/8/4 23:59
 * @version     v1.0
 * @filename    server
 * @description
 ***************************************************************************/
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type FoodService struct {
}

func (f *FoodService) SayName(request string, resp *string) error {
	*resp = "您点的菜是：" + request
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 解析json
	}
}
