package main

import (
	"fmt"
	"log"

	pb "github.com/nXnUs25/gRPCHellow/proto"
)

func (s *GreetingsServer) GreetingsManyTimees(in *pb.GreetingsRequest, stream pb.GreetingsService_GreetingsManyTimeesServer) error {
	log.Printf("GreetingsManyTimees executed: %v", in)

	for x := 1; x <= 10; x++ {
		res := fmt.Sprintf("Hello %v grettings for %v time.", in.Fname, x)

		stream.Send(
			&pb.GreetingsResponse{
				Answer: res,
			},
		)
	}
	return nil
}
