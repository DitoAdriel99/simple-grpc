package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Sum func running: %v\n", req)
	return &pb.SumResponse{
		Result: req.FirstNumber + req.SecondNumber,
	}, nil
}
