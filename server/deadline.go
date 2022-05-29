package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nXnUs25/gRPCHellow/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GreetingsServer) GreetingsWithDeadline(cxt context.Context, in *pb.GreetingsRequest) (*pb.GreetingsResponse, error) {
	log.Printf("GreetingsWithDeadline func called %v\n", in)

	for i := 1; i < 3; i++ {
		if cxt.Err() == context.DeadlineExceeded {
			log.Println("Client cancelled the request")
			return nil, status.Error(codes.Canceled, "Client cancelled the request")
		}
		time.Sleep(2 * time.Second)
	}
	return &pb.GreetingsResponse{Answer: "Greetings with deadline from Hello grpc app " + in.Fname}, nil
}
