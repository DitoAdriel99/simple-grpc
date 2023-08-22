package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("create blog was invoked")

	blog := &pb.Blog{
		AuthorId: "Dito",
		Title:    "this is title",
		Content:  "this is content",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("err while create: %v\n", err)
	}

	log.Printf("blog has been created: %s\n", res.Id)
	return res.Id

}
