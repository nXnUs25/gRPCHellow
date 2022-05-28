package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

func (s *GreetingsServer) GreetingEveryone(stream pb.GreetingsService_GreetingEveryoneServer) error {
	log.Println("Everyone GHreetings executed")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading stream: \n%v", err)
		}
		res := fmt.Sprintf("Hello %s !\n", req.GetFname())

		err = stream.Send(&pb.GreetingsResponse{
			Answer: res,
		})

		if err != nil {
			log.Fatalf("Error while sending response to client: \n%v", err)
		}
	}
}
