package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Print("Do max was invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while openign stream: %v\n", err)
	}
	waitc := make(chan struct{})
	numbers := []int32{4, 7, 2, 19, 4, 6, 32}
	for _, number := range numbers {
		log.Printf("Sending Number: %d\n", number)
		stream.Send(&pb.MaxRequest{Number: number})
		time.Sleep(1 * time.Second)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Problem whil reading stream: %v\n", err)
			break
		}
		log.Printf("Receiving a new max number: %d\n", res.Result)
	}
	stream.CloseSend()

	close(waitc)
	<-waitc
}
