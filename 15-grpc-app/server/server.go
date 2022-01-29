package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"net"
	"time"

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

func (s *server) GeneratePrimes(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Request receive for generating prime from %d to %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			time.Sleep(500 * time.Millisecond)
			res := &proto.PrimeResponse{
				PrimeNumber: no,
			}
			fmt.Println("Sending prime no : ", no)
			stream.Send(res)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	if no <= 0 {
		return false
	}
	if no <= 3 {
		return true
	}
	for i := int32(2); i <= no-1; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
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
