package main

import (
	"context"
	"log"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

func (s *GreetingsServer) Greetings(cxt context.Context, in *pb.GreetingsRequest) (*pb.GreetingsResponse, error) {
	log.Printf("Greetings func called %v\n", in)

	return &pb.GreetingsResponse{Answer: "Greetings from Hellow grpc app " + in.Fname}, nil
}
