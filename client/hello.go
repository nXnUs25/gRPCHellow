package main

import (
	"context"
	"log"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

func Greetings(c pb.GreetingsServiceClient) {
	res, err := c.Greetings(context.Background(), &pb.GreetingsRequest{
		Fname: "Auggie",
	})
	if err != nil {
		log.Fatalf("Failed to greeting: %v", err)
	}

	log.Printf("%v", res.Answer)
}
