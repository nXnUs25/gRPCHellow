package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/nXnUs25/gRPCHellow/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s - %v", addr, err)
	}
	defer listen.Close()

	log.Printf("listening on %s\n", addr)

	opts := []grpc.ServerOption{}

	tls := true // change that to true if needed
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	defer s.Stop()

	pb.RegisterGreetingsServiceServer(s, &GreetingsServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve on %s - %v", addr, err)
	}

	sChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		<-sChan
		done <- true
	}()
	<-done
	log.Println("Sopping GRPC Server.")
	s.GracefulStop()
}
