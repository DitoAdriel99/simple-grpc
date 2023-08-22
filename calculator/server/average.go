package main

import (
	pb "github.com/DitoAdriel99/grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("average was invoke ")
	var sum int32 = 0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("error whilew read client stream : %v\n", err)
		}
		log.Printf("Receiving %v\n", req.Number)
		sum += req.Number
		count++
	}
}
