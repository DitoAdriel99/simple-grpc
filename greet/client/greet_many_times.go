package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("dogrett many times")

	req := &pb.GreetRequest{FirstName: "Dito"}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greetmanytimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("GreetManytimes: %s\n", msg.Result)
	}
}
