package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

func (s *GreetingsServer) LongBackGretings(stream pb.GreetingsService_LongBackGretingsServer) error {
	log.Println("LongBack for GHreetings executed")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetingsResponse{
				Answer: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		res += fmt.Sprintf("Hello %s\n", req.GetFname())
	}

}
