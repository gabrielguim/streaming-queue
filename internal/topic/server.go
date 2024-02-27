package topic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"slices"
	"streaming-queue/proto"
)

var (
	topics = &[]Topic{}
)

type Server struct {
	proto.UnimplementedTopicServer
}

func (s *Server) Create(ctx context.Context, topic *proto.TopicRequest) (*proto.Void, error) {
	*topics = append(*topics, *CreateTopic(topic.Name))

	return &proto.Void{}, nil
}

func (s *Server) Delete(ctx context.Context, topic *proto.TopicRequest) (*proto.Void, error) {
	*topics = slices.DeleteFunc(*topics, func(t Topic) bool {
		return t.name == topic.Name
	})

	return &proto.Void{}, nil
}

func (s *Server) Publish(ctx context.Context, message *proto.Message) (*proto.Void, error) {
	for _, topic := range *topics {
		if message.Topic == topic.name {
			topic.Publish(message)
			break
		}
	}

	go WriteMessage(message)

	return &proto.Void{}, nil
}

func streamMessage(stream proto.Topic_SubscribeServer, message *proto.Message) {
	err := stream.Send(message)
	if err != nil {
		fmt.Printf("Stream ended: %v\n", err)
		panic(err)
	}
}

func (s *Server) Subscribe(request *proto.SubscribeRequest, stream proto.Topic_SubscribeServer) error {
	var topic Topic
	for _, t := range *topics {
		if request.Topic == t.name {
			topic = t
			break
		}
	}

	if !Exists(topic) {
		return status.Errorf(codes.NotFound, "Topic not found")
	}

	if len(request.Offset) > 0 {
		messages := ReadMessages(request.Offset, request.Topic)

		for _, message := range messages {
			streamMessage(stream, &message)
		}
	}

	observer := &Observer{
		func(message *proto.Message) {
			streamMessage(stream, message)
		},
		request.Consumer,
	}

	topic.RegisterObserver(observer)

	for {
		select {
		case <-stream.Context().Done():
			topic.UnregisterObserver(observer)
			return status.Errorf(codes.Canceled, "Stream ended")
		}
	}
}
