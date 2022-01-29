package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewAppServiceClient(clientConn)

	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	ctx := context.Background()
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Adding 100 and 200, result = %d\n", res.GetResult())
}
