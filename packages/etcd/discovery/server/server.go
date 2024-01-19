package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"learn-go/packages/etcd/discovery"
	pb "learn-go/packages/etcd/discovery/proto"
	"net"
	"strconv"
)

var (
	port = flag.Int("port", 50052, "")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("Recv Client msg: %v\n", in.Msg)
	return &pb.HelloReply{
		Msg: "Hello client",
	}, nil
}

func serverRegister(s grpc.ServiceRegistrar, srv pb.GreeterServer) {
	pb.RegisterGreeterServer(s, srv)
	s1 := &discovery.Service{
		Name:     "hello.Greeter",
		Port:     strconv.Itoa(*port),
		IP:       "127.0.0.1",
		Protocol: "grpc",
	}
	go discovery.ServiceRegister(s1)
}

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	serverRegister(s, &server{})
	if err = s.Serve(listen); err != nil {
		panic(err)
	}
}
