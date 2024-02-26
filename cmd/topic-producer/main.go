package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"streaming-queue/proto"
	"time"
)

func sendMessage(grpcClient proto.TopicClient, topic string) {
	date := time.Now().Format(time.RFC3339)
	message := &proto.Message{
		Id:        rand.Uint32(),
		Topic:     topic,
		CreatedAt: date,
		Title:     fmt.Sprintf("Message '%d'", rand.Int()),
	}

	_, err := grpcClient.Publish(context.Background(), message)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Sent message: %v\n", message)
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

	go func() {
		topic := "topic-1"
		_, err = grpcClient.Create(context.Background(), &proto.TopicRequest{
			Name: topic,
		})

		if err != nil {
			panic(err)
		}

		for {
			time.Sleep(1 * time.Second)
			sendMessage(grpcClient, topic)
		}
	}()

	go func() {
		topic := "topic-2"
		_, err = grpcClient.Create(context.Background(), &proto.TopicRequest{
			Name: topic,
		})

		if err != nil {
			panic(err)
		}

		for {
			time.Sleep(2 * time.Second)
			sendMessage(grpcClient, topic)
		}
	}()

	wait := make(chan bool)
	<-wait

}
