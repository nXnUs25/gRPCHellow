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
	opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(), CheckHeaderInterceptor()))
	s := grpc.NewServer(opts...)
	defer s.Stop()

	sChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		sig := <-sChan
		log.Printf("Got signal [%v]", sig)
		log.Println("Stopping gRPCHellow server...")
		s.GracefulStop()
		done <- true
	}()

	pb.RegisterGreetingsServiceServer(s, &GreetingsServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve on %s - %v", addr, err)
	}
	<-done
	log.Println("Stopped GRPC Server.")
}
