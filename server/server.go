// protoc *.proto --go_out=./ --go-grpc_out=./
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sdejesusp/golang-grpc-basic1/pb/pb"
	"google.golang.org/grpc"
)


type server struct {
	pb.GreetingServiceServer
}

func (s *server) Greeting(ctx context.Context, req *pb.GreetingServiceRequest) (*pb.GreetingServiceReply, error) {
	return &pb.GreetingServiceReply {
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}