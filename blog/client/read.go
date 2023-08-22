package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"log"
)

func read(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("read blog was invoked ")

	req := &pb.BlogId{
		Id: id,
	}

	obj, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("error read blog: %v\n", err)
	}

	log.Println("Blog was read: %v\n", obj)
	return obj
}
