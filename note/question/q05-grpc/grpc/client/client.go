/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/8/5 22:58
 * @version     v1.0
 * @filename    client.go
 * @description
 ***************************************************************************/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn-go/note/question/q05-grpc/grpc/proto/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewStudyClient(conn)
	res, err := client.Study(context.Background(), &pb.BookRequest{Name: "数学"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Msg)
}
