package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	fmt.Println("Sending result :", result)
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	var s = &server{}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, s)
	grpcServer.Serve(listener)
}
