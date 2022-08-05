/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/8/5 23:44
 * @version     v1.0
 * @filename    client.go
 * @description
 ***************************************************************************/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn-go/note/question/q05-grpc/stream/proto/pb"
	"sync"
)

func main() {
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewFoodServiceClient(conn)
	//// 服务端流模式
	//res, err := client.SayName(context.Background(), &pb.FoodStreamRequest{
	//	Name: "麻辣小龙虾",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//for {
	//	recv, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(recv.Msg)
	//}

	//// 客户端流模式
	//ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	//defer cancel()
	//postNameClient, err := client.PostName(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//foods := []string{"东坡肘子", "回锅肉", "夫妻肺片", "水煮牛肉"}
	//for _, item := range foods {
	//	fmt.Println("上菜：" + item)
	//	err = postNameClient.Send(&pb.FoodStreamRequest{Name: item})
	//	time.Sleep(time.Second)
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//}

	// 双向流模式
	fullClient, err := client.FullStream(context.Background())
	if err != nil {
		panic(err)
	}
	foods := []string{"东坡肘子", "回锅肉", "夫妻肺片", "水煮牛肉"}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			item, err2 := fullClient.Recv()
			if err2 != nil {
				fmt.Println(err)
			}
			fmt.Println(item.Msg)
		}
	}()
	go func(s []string) {
		defer wg.Done()
		for _, item := range s {
			err := fullClient.Send(&pb.FoodStreamRequest{Name: item})
			if err != nil {
				fmt.Println(err)
			}
		}
	}(foods)
	wg.Wait()
}
