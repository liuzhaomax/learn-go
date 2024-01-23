package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"learn-go/note/question/q06-grpc_advanced/grpc_interceptor/proto/pb"
	"net"
	"time"
)

type ToDo struct {
}

func (t *ToDo) DoWork(ctx context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("no metadata")
	}
	fmt.Println(md["name"][0])
	for k, v := range md {
		fmt.Printf("%s:%s\n", k, v)
	}
	time.Sleep(time.Second * 2)
	fmt.Println(req.Todo + "已完成！")
	return &pb.TodoResponse{Done: true}, nil
}

func MyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	now := time.Now()
	resp, err = handler(ctx, req)
	d := time.Now().Sub(now)
	fmt.Printf("执行时间: %d\n", d.Milliseconds())
	// st := status.New(codes.Unauthenticated, "未授权")
	// err = st.Err()
	err = status.Error(codes.Unauthenticated, "未授权啊")
	fmt.Println(int(codes.Unauthenticated))
	return
}

func main() {
	serverOption := grpc.UnaryInterceptor(MyInterceptor)
	server := grpc.NewServer(serverOption)
	pb.RegisterToDoServiceServer(server, &ToDo{})
	listen, err := net.Listen("tcp", "0.0.0.0:9096")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
