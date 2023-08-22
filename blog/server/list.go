package main

import (
	"context"
	"fmt"
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListBlog(in *pb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("list was invoked")
	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal, "unknown internal error",
		)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("error nya : %v\n", err))
		}
		stream.Send(documentTobBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown error: %v\n", err))
	}
	return nil
}
