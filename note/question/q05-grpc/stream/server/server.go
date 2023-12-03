package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learn-go/note/question/q05-grpc/stream/proto/pb"
	"net"
	"sync"
	"time"
)

type FoodInfo struct {
}

func (f *FoodInfo) SayName(request *pb.FoodStreamRequest, server pb.FoodService_SayNameServer) error {
	fmt.Println("SayName req came")
	err := server.Send(&pb.FoodStreamResponse{
		Msg: "您点的是：" + request.Name,
	})
	if err != nil {
		return err
	}
	return nil
}

func (f *FoodInfo) PostName(server pb.FoodService_PostNameServer) error {
	fmt.Println("PostName req came")
	for {
		recv, err := server.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("您请慢用：" + recv.Name)
	}
	return nil
}

func (f *FoodInfo) FullStream(server pb.FoodService_FullStreamServer) error {
	fmt.Println("FullStream req came")
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan string, 5)
	go func() {
		defer wg.Done()
		for {
			item, err := server.Recv()
			if item == nil {
				return
			}
			if err != nil {
				fmt.Println(err)
			}
			ch <- item.Name
			fmt.Println("已下单: " + item.Name)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			foodName := <-ch
			err := server.Send(&pb.FoodStreamResponse{
				Msg: "菜" + foodName + "做好了",
			})
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9092")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterFoodServiceServer(server, &FoodInfo{})
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
