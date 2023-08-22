package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do greet")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Dito",
	})
	if err != nil {
		log.Fatal("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res)
}
