package main

import (
	"context"
	"fmt"
	"grpc/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello from the Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v:", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Start the Unary RPC")
	req := &calculatorpb.SumRequest{
		FirstNumber:  -3,
		SecondNumber: 40.8,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v:", err)
	}
	log.Printf("The total sum is: %v", res.SumResult)
}
