package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"log"
)

func update(c pb.BlogServiceClient, id string) {
	log.Printf("update was invoked")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Yeyen",
		Title:    "hai",
		Content:  "aku sayang kamu",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("err update : %v\n", err)
	}

	log.Println("blog was updated")
}
