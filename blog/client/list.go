package main

import (
	"context"
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"io"
	"log"
)

func list(c pb.BlogServiceClient) {
	log.Println("list blog was invoked")
	stream, err := c.ListBlog(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("error whil calling list: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error whil reacing stream: %v\n", err)
		}
		log.Printf("response data: %v\n", res)
	}

}
