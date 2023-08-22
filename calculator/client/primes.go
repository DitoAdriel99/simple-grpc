package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimeRequest{
		Number: 12390392840,
	}
	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while stream primes: %v\n")
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while readind stream: %v\n")
		}

		log.Printf("Primes: %d", res.Result)
	}
}
