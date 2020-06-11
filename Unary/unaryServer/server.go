package main

import (
	"context"
	"fmt"
	"go-gRPC-examples/Unary/unarypb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Addition(ctx context.Context, req *unarypb.SumRequest) (*unarypb.SumResponse, error) {
	fmt.Println("Recieved Request ..", req)
	firstNumber := req.FirstNo
	secondNumber := req.SecondNo
	sum := firstNumber + secondNumber
	res := &unarypb.SumResponse{
		Result: sum,
	}
	return res, nil
}
func main() {
	fmt.Println("Unary Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	unarypb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
