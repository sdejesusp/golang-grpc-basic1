package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sdejesusp/golang-grpc-basic1/pb/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8081", opts)
	if err != nil{
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewGreetingServiceClient(cc)
	request := &pb.GreetingServiceRequest{Name: "Gophers"}

	resp, err := client.Greeting(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", resp.Message)

}