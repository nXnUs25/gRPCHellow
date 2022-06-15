package main

import (
	"context"
	"testing"

	pb "github.com/nXnUs25/gRPCHellow/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestLongGreet(t *testing.T) {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetingsServiceClient(conn)

	requests := []*pb.GreetingsRequest{
		{
			Fname: "Auggie",
		},
		{
			Fname: "Marie",
		},
		{
			Fname: "Test",
		},
	}

	stream, err := c.LongBackGretings(context.Background())

	if err != nil {
		t.Errorf("GreetingsManyTimes(%v) got unexpected error", err)
	}

	for _, req := range requests {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		t.Errorf("Error while receiving response from LongGreet: %v", err)
	}

	expected := "Hello Auggie\nHello Marie\nHello Test\n"

	if res.GetAnswer() != expected {
		t.Errorf("Expected \"%s\" elements, got: \"%v\"", expected, res.GetAnswer())
	}
}
