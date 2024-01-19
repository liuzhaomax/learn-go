package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learn-go/packages/etcd/discovery"
	pb "learn-go/packages/etcd/discovery/proto"
	"log"
	"time"
)

func getServerAddr(svcName string) string {
	s := discovery.ServiceDiscover(svcName)
	if s == nil {
		return ""
	}
	if s.IP == "" && s.Port == "" {
		return ""
	}
	return s.IP + ":" + s.Port
}

func sayHello() {
	addr := getServerAddr("hello.Greeter")
	if addr == "" {
		log.Println("未发现可用服务")
		return
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	in := &pb.HelloRequest{
		Msg: "Hello server",
	}

	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), in)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Recv server msg: %v\n", r.Msg)
}

func main() {
	go discovery.WatchServiceName("hello.Greeter")
	for {
		sayHello()
		time.Sleep(time.Second * 2)
	}
}
