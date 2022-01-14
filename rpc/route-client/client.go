/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/13 20:56
 * @version     v1.0
 * @filename    client.go
 * @description
 ***************************************************************************/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "maxmicro/test/rpc/route"
)

func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client net error")
		return
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	runFirst(client)
}
