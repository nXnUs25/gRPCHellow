package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nXnUs25/gRPCHellow/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GreetingsDeadlien(c pb.GreetingsServiceClient, sec int32) {
	timeout := time.Duration(sec) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetingsWithDeadline(ctx, &pb.GreetingsRequest{
		Fname: "Auggie",
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Recived error message from server %v", e.Message())
			log.Printf("Received error code status %v", e.Code())

			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded")
				return
			}
		} else {
			log.Fatalf("Non gRPC error: Failed to greeting: %v", err)
		}

	}

	log.Printf("%v", res.Answer)
}
