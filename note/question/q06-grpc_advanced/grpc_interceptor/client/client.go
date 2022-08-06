/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/8/6 21:42
 * @version     v1.0
 * @filename    client.go
 * @description
 ***************************************************************************/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	client.DoWork(context.Background(), &pb.TodoRequest{Todo: "Go语言"})
}
