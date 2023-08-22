package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/calculator/proto"
	"log"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("do avg ni")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error while opening the stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)
		stream.Send(&pb.AverageRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response : %v\n", err)
	}
	log.Printf("Avg: %f\n", res.Result)
}
