package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/DavidHODs/gRPC/protocol/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {}

func (*server) Add(ctx context.Context, req *calculatorpb.AdditionRequest) (*calculatorpb.AdditionResponse, error) {
	fmt.Printf("Add function was invoked with %v\n", req)
	firstNumber := req.GetAddition().GetFirstInteger()
	secondNumber := req.GetAddition().GetSecondInteger()
	result := firstNumber + secondNumber

	res := &calculatorpb.AdditionResponse {
		Result: result,
	}
	return res, nil
}

func (*server) AddManyTimes(req *calculatorpb.AddManyTimesRequest, stream calculatorpb.AdditionService_AddManyTimesServer) error {
	fmt.Printf("AddManyTimes function was invoked with %v\n", req)
	firstInteger := req.Addition.GetFirstInteger()
	secondInteger := req.Addition.GetSecondInteger()
	result := firstInteger + secondInteger
	
	var i int32
	for i = 0; i < 100; i++ {
		res := &calculatorpb.AddManyTimesResponse{
			Result: i + result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterAdditionServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}