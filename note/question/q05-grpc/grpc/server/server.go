package main

import (
	"context"
	"google.golang.org/grpc"
	"learn-go/note/question/q05-grpc/grpc/proto/pb"
	"net"
)

type BookInfo struct {
}

func (b *BookInfo) Study(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{
		Msg: "我要学习: " + req.Name,
	}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterStudyServer(server, &BookInfo{})
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
