package main

import (
	"fmt"
	"go-gRPC-examples/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello this is a client code")

	// Create a connection to a server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	fmt.Printf("Created client %f", c)
}
