package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"streaming-queue/proto"
)

func startConsumer(grpcClient proto.TopicClient, consumer string) {
	request := &proto.SubscribeRequest{
		Topic:    "my-topic",
		Consumer: consumer,
	}

	stream, err := grpcClient.Subscribe(context.Background(), request)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for {
		message, err := stream.Recv()

		if err != nil {
			fmt.Printf("Error while reading: %v\n", err)
			stream.Context().Done()
		} else {
			fmt.Printf("Receive message '%v' in consumer '%s'\n", message.String(), consumer)
		}
	}

}

func main() {
	var opts []grpc.DialOption
	serverAddr := "localhost:9000"

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(serverAddr, opts...)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	grpcClient := proto.NewTopicClient(conn)

	go startConsumer(grpcClient, "consumer 1")
	go startConsumer(grpcClient, "consumer 2")

	wait := make(chan bool)
	<-wait
}
