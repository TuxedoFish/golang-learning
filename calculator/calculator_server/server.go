package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/TuxedoFish/golang-learning/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with: %v \n", req)
	result := req.GetA() + req.GetB()
	res := calculatorpb.SumResponse{
		Result: result,
	}
	return &res, nil
}

func main() {
	fmt.Println("Starting calculator service")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
