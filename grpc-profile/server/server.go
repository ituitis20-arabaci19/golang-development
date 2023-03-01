package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ituitis20-arabaci19/golang-development-template/grpc-profile/handler"
	"github.com/ituitis20-arabaci19/golang-development-template/grpc-profile/proto/profile"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	profile.RegisterProfileServiceServer(s, &handler.Profile{Id: 0, Persons: make(map[string]handler.User)})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
