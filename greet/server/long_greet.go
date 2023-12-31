package main

import (
	"fmt"
	pb "github.com/DitoAdriel99/grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Longgreet function was invoked")
	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("error while reading clien strwam : %v\n", err)
		}
		log.Printf("Receiving: %v\n", req)

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
