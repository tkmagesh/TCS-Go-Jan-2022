package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	//doRequestResponse(ctx, client)
	//doServerStreaming(ctx, client)
	//doClientStreaming(ctx, client)
	doBidirectionalStreaming(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}

	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Adding 100 and 200, result = %d\n", res.GetResult())
}

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	stream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		primeNo := res.GetPrimeNumber()
		fmt.Println("Prime no received : ", primeNo)
	}
}

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	var nos = []int32{5, 2, 4, 3, 8, 7, 9, 6, 1}
	stream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Sending Number : ", no)
		req := &proto.AverageRequest{
			No: no,
		}
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average Received : ", res.GetAverage())
}

func doBidirectionalStreaming(ctx context.Context, client proto.AppServiceClient) {
	stream, err := client.GenerateMultiPrimes(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()

			if err != nil {
				log.Fatalln(err)
			}

			primeNo := res.GetPrimeNumber()
			fmt.Println("Prime Number received: ", primeNo)
		}
		done <- struct{}{}
	}()

	req1 := &proto.PrimeRequest{
		Start: 3,
		End:   30,
	}
	fmt.Println("Sending request - 1")
	stream.Send(req1)

	time.Sleep(2 * time.Second)

	req2 := &proto.PrimeRequest{
		Start: 31,
		End:   60,
	}
	fmt.Println("Sending request - 2")
	stream.Send(req2)

	time.Sleep(2 * time.Second)

	req3 := &proto.PrimeRequest{
		Start: 61,
		End:   100,
	}
	fmt.Println("Sending request - 3")
	stream.Send(req3)

	<-done
}
