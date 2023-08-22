package main

import (
	"context"
	"fmt"
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*pb.Empty, error) {
	log.Println("Delete was invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error id: %v\n", err))
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error delete: %v\n", err))
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("error delete: %v\n", err))
	}
	return &pb.Empty{}, nil
}
