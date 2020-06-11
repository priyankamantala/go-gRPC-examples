package main

import (
	"context"
	"fmt"
	"go-gRPC-examples/Unary/unarypb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Unary Client")

	// Create a connection to a server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect", err)
	}
	defer conn.Close()
	c := unarypb.NewCalculatorServiceClient(conn)

	unaryRPC(c)
}

func unaryRPC(c unarypb.CalculatorServiceClient) {
	fmt.Println("Unary Addition")
	request := &unarypb.SumRequest{
		FirstNo:  10,
		SecondNo: 20,
	}
	response, err := c.Addition(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling addition RPC %v", err)
	}
	fmt.Println("Result:", response.Result)
}
