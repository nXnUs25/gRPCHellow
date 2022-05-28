package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

func GretingsEveryone(c pb.GreetingsServiceClient) {
	log.Println("GreetingEveryone")

	regs := []*pb.GreetingsRequest{
		{Fname: "Adam"},
		{Fname: "John"},
		{Fname: "John Smith"},
	}
	stream, err := c.GreetingEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongBackGretings: %v", err)
	}

	done := make(chan struct{})

	go func() {
		for _, req := range regs {
			log.Printf("Sending req %v\n", req)
			stream.Send(req)
			//time.Sleep(2 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while closing stream: %v", err)
				break
			}
			log.Printf("Greetings to Everyone: %v", res.GetAnswer())
		}
		close(done)
	}()
	<-done
}
