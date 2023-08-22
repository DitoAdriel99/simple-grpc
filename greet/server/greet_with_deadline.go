package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greetwithdeadline was invoked: %v\n", in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client cancel the request")
			return nil, status.Error(codes.Canceled, "The Client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Helo" + in.FirstName,
	}, nil
}
