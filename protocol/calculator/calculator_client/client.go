package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/DavidHODs/gRPC/protocol/calculator/calculatorpb"
	"google.golang.org/grpc"
)


func main () {
	opts := grpc.WithInsecure()

	conn, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewAdditionServiceClient(conn)

	// doUnary(c)
	doServerStreaming(c)
}

// func doUnary(c calculatorpb.AdditionServiceClient) {
// 	fmt.Println("Starting to do a Unary RPC...")
// 	req := &calculatorpb.AdditionRequest{
// 		Addition: &calculatorpb.Addition{
// 			FirstInteger: 10,
// 			SecondInteger: 3,
// 		},
// 	}

// 	res, err := c.Add(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error while calling Add RPC: %v", err)
// 	}
// 	log.Printf("Response from Add: %v", res.Result)
// }

func doServerStreaming(c calculatorpb.AdditionServiceClient) {
	fmt.Println("Starting to do a Streaming RPC...")

	req := &calculatorpb.AddManyTimesRequest{
		Addition: &calculatorpb.Addition{
			FirstInteger: 2,
			SecondInteger: 3,
		},
	}
	resStream, err := c.AddManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling AddManyTimes RPC: %v", err)
	}
	for { msg, err := resStream.Recv() 
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from AddManyTimes: %v", msg.GetResult())
	} 
}