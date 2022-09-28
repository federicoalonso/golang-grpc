package main

import (
	"log"
	"net"

	"golang-grpc.com/packages/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
}
