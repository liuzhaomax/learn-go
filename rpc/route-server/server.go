/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/13 20:55
 * @version     v1.0
 * @filename    server.go
 * @description
 ***************************************************************************/
package main

import (
	"context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	pb "learn-go/rpc/route"
	"log"
	"net"
)

type routeGuideServer struct {
	features []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

// unary
func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, nil
}

// server side streaming
func (s *routeGuideServer) ListFeature(rectangle *pb.Rectangle, stream pb.RouteGuide_ListFeatureServer) error {
	return nil
}

// client side streaming
func (s *routeGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}

// bidirectional streaming
func (s *routeGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}

func newServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
			{
				Name: "shanghai jiaotong No 888",
				Location: &pb.Point{
					Latitude:  310235000,
					Longitude: 121437403,
				},
			}, {
				Name: "fudan",
				Location: &pb.Point{
					Latitude:  312978870,
					Longitude: 121503457,
				},
			}, {
				Name: "huadong",
				Location: &pb.Point{
					Latitude:  311416130,
					Longitude: 121424904,
				},
			},
		},
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("net error")
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("net error")
		return
	}
}
