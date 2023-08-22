package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("do sqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("error message from server: %s\n", e.Message())
			log.Printf("error message from server: %s\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("We Propably sent a negative number")
				return
			}
		} else {
			log.Fatalf("a Non GRPC error: %v\n", err)
		}
	}
	log.Printf("Sqer: %f\n", res.Result)
}
