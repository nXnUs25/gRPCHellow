package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

func LongBackGretings(c pb.GreetingsServiceClient) {
	log.Println("LongBackGretings")

	regs := []*pb.GreetingsRequest{
		{Fname: "Adam"},
		{Fname: "John"},
		{Fname: "John Smith"},
	}
	stream, err := c.LongBackGretings(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongBackGretings: %v", err)
	}
	for _, req := range regs {
		log.Printf("Sending req %v\n", req)
		stream.Send(req)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while closing stream: %v", err)
	}
	log.Printf("Long Greetings: \n%v\n", res.GetAnswer())
}
