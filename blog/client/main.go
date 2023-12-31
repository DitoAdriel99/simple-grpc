package main

import (
	pb "github.com/DitoAdriel99/grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}
	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	id := createBlog(c)
	read(c, id)
	list(c)
	//update(c, id)
	//read(c, "sdsds")
	delete(c, id)
	list(c)

}
