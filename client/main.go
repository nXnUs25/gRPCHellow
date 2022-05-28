package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

var addr string = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewGreetingsServiceClient(conn)
	Greetings(c)
	ManyGreetings(c)
	LongBackGretings(c)
}
