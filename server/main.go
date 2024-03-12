package main

import (
	"log"
	"net"

	pb "grpc-platform/pb/platform"

	"google.golang.org/grpc"
	// mongo "tutorial/grpc-platform/server/db"
)

const (
	// Port for gRPC server to listen to
	PORT = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterPlatformServiceServer(s, &PlatformServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
