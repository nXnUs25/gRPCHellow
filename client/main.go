package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change that to true if needed
	opts := []grpc.DialOption{}

	//conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if tls {
		certFile := "ssl/server.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewGreetingsServiceClient(conn)
	Greetings(c)
	ManyGreetings(c)
	LongBackGretings(c)
	GretingsEveryone(c)
	GreetingsDeadlien(c, 3)
	GreetingsDeadlien(c, 10)
}
