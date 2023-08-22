package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("do long greet ")
	reqs := []*pb.GreetRequest{
		{FirstName: "dito"},
		{FirstName: "Yeyen"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling longgreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving longgreet %v\n", err)
	}

	log.Printf("Longgreet: %s\n", res.Result)
}
