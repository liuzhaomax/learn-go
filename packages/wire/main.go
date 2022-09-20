package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"learn-go/packages/wire/pb"
	"net"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "输入ip")
	port := flag.Int("port", 9095, "输入port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	injector, _, _ := InitInjector()
	server := grpc.NewServer()
	pb.RegisterDataServiceServer(server, injector.Service) // 这里是重点，将业务逻辑注入server
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
