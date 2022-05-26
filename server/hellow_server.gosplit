package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nXnUs25/gRPCHellow/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type GreetingsServer struct {
	pb.GreetingsServiceServer
}

func (s *GreetingsServer) Greetings(cxt context.Context, in *pb.GreetingsRequest) (*pb.GreetingsResponse, error) {
	log.Printf("Greetings func called %v\n", in)

	return &pb.GreetingsResponse{Answer: "Greetings from Hellow grpc app " + in.Fname}, nil
}

func main() {
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %s - %v", addr, err)
	}

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetingsServiceServer(s, &GreetingsServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve on %s - %v", addr, err)
	}
}
