package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TuxedoFish/golang-learning/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Hello I'm a calculator client \n")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewSumServiceClient(cc)

	doSum(c)
}

func doSum(c calculatorpb.SumServiceClient) {
	req := calculatorpb.SumRequest{
		A: 3,
		B: 10,
	}

	fmt.Printf("Request to sum: %[1]v + %[2]v \n", 3, 10)
	res, err := c.Sum(context.Background(), &req)

	if err != nil {
		log.Fatalf("Failed to sum: %v", err)
	}

	fmt.Printf("Response: %v \n", res.Result)
}
