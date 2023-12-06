package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learn-go/note/question/q06-grpc_advanced/grpc_interceptor/proto/pb"
	"time"
)

func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	now := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	d := time.Now().Sub(now)
	fmt.Printf("客户端执行时间 %d \n", d.Milliseconds())
	return err
}

func main() {
	opt := grpc.WithUnaryInterceptor(clientInterceptor)
	conn, err := grpc.Dial("127.0.0.1:9096", grpc.WithInsecure(), opt)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := pb.NewToDoServiceClient(conn)

	// add metadata
	//md1 := metadata.New(map[string]string{
	//	"name": "metadata",
	//})
	md2 := metadata.Pairs("name", "%3#_@")
	ctx := metadata.NewOutgoingContext(context.Background(), md2)

	res, err := client.DoWork(ctx, &pb.TodoRequest{Todo: "Go语言"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Done)
}
