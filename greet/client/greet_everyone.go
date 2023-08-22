package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("Do greet everyone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream: %v\n")
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Dito"},
		{FirstName: "Yeyen"},
		{FirstName: "Budi"},
		{FirstName: "Yanto"},
	}
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send Request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error whil receiving stream: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res)
		}
		close(waitc)
	}()
	<-waitc
}
