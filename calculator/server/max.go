package main

import (
	pb "github.com/DitoAdriel99/grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("max was invoked")
	var maximum int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error While reading stream: %v\n", err)
		}
		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{Result: maximum})
			if err != nil {
				log.Fatalf("Error While send stream: %v\n", err)

			}
		}
	}
	return nil
}
