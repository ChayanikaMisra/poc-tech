package main

import (
	"context"
	"fmt"
	"golang-grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.HelloWorldServiceServer
}

func (s *server) Greeting(ctx context.Context, req *pb.HelloWorldServiceRequest) (*pb.HelloWorldServiceReply, error) {
	return &pb.HelloWorldServiceReply{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func main() {
	fmt.Println("start grpc server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}