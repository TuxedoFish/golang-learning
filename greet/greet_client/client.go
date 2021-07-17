package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TuxedoFish/golang-learning/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Create client: %f", c)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Printf("Starting to do Unary request")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Harry",
			LastName:  "Liversedge",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to call greet: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
