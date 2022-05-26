package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

func ManyGreetings(c pb.GreetingsServiceClient) {
	stream, err := c.GreetingsManyTimees(context.Background(), &pb.GreetingsRequest{
		Fname: "Auggie",
	})
	if err != nil {
		log.Fatalf("Failed to greeting: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read stream on %s - %v", res, err)
		}
		log.Printf("%v", res.Answer)
	}

}
