package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/DavidHODs/gRPC/protocol/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()

	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("Created new client: %f", c)

	// doUnary(c)
	
	// doServerStreaming(c)

	// doClientStreaming(c)

	doBidiStreaming(c)
	
}

// func doUnary(c greetpb.GreetServiceClient) {
// 	fmt.Println("Starting to do a Unary RPC...")
// 	req := &greetpb.GreetRequest{
// 		Greeting: &greetpb.Greeting{
// 			FirstName: "David",
// 			LastName: "Oluwatobi",
// 		},
// 	}

// 	res, err := c.Greet(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error while calling Greet RPC: %v", err)
// 	}
// 	log.Printf("Response from Greet: %v", res.Result)
// }


// func doServerStreaming(c greetpb.GreetServiceClient) {
// 	fmt.Println("Starting to do a Streaming RPC...")

// 	req := &greetpb.GreetManyTimesRequest{
// 		Greeting: &greetpb.Greeting{
// 			FirstName: "Klaus",
// 			LastName: "Stefan",
// 		},
// 	}

// 	resStream, err := c.GreetManyTimes(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
// 	}
// 	for {
// 		msg, err := resStream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Error while reading stream: %v", err)
// 		}
// 		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
// 	}
// }

// func doClientStreaming(c greetpb.GreetServiceClient) {
// 	fmt.Println("Starting to do a Client Streaming RPC...")

// 	requests := []*greetpb.LongGreetRequest{
// 		{
// 			Greeting: &greetpb.Greeting{FirstName: "Stefan"},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Isaih",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Luke",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Klaus",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Elijah",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Kanmi",
// 			},
// 		},
// 	}

// 	stream, err := c.LongGreet(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error while calling LongStreet: %v", err)
// 	}

// 	for _, req := range requests {
// 		fmt.Printf("Sending req: %v\n", req)
// 		stream.Send(req)
// 		time.Sleep(100 * time.Millisecond)
// 	}

// 	res, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatalf("Error while receiving response from LongGreet: %v", err)
// 	}
// 	fmt.Printf("LongGreet Response: %v\n", res)
// }

func doBidiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do bi-directional streaming")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return 
	}

	requests := []*greetpb.GreetEveryOneRequest{
		{
			Greeting: &greetpb.Greeting{FirstName: "Stefan"},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Isaih",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Luke",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Klaus",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Elijah",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Kanmi",
			},
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
	} ()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break;
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break;
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	} ()

	<-waitc
}