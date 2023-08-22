package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"log"
)

func delete(c pb.BlogServiceClient, in string) {
	log.Println("delete was invoked")

	req := &pb.BlogId{Id: in}

	_, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Printf("errorr delte: %v\n", err)
	}

	log.Println("delete was success")
}
