package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"streaming-queue/internal/topic"
	"streaming-queue/proto"
)

func main() {
	println("Running gRPC topic-server")

	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTopicServer(grpcServer, &topic.Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
