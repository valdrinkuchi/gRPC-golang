package main

import (
	"context"
	"fmt"
	"grpc/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello from the Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v:", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Start the Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Valdrin",
			LastName:  "Kuchi",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v:", err)
	}
	log.Printf("Response from Greet %v", res.Result)
}
