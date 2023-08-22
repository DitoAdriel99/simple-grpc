package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("do greet with deadline was invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Dito"}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
			} else {
				log.Fatalf("Grpc error: %v\n", err)
			}
		} else {
			log.Fatalf("A non grpc error: %v\n", err)
		}
	}

	log.Printf("Greet with deadline: %s\n", res.Result)
}
