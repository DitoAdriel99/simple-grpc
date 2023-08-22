package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/calculator/proto"
	"log"
)

func doCalculator(c pb.CalculatorServiceClient) {
	log.Println("do calculator here")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatal("Could not calculated: %v\n", err)
	}
	log.Printf("Result: %s\n", res)
}
