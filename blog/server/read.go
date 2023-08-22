package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("read blog was invoked with %v\n", in.Id)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		log.Fatalf("error objec from hex : %v\n", err)
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}
	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err = res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find id",
		)
	}

	return documentTobBlog(data), nil
}
