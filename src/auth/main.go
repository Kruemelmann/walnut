package main

import (
	"fmt"
	"log"
	"net"

	walnut "github.com/kruemelmann/walnut/src/auth/genproto"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting walnut auth service")
	lis, err := net.Listen("tcp", fmt.Sprintf(":50051"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	walnut.RegisterAuthServiceServer(s, &authServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
