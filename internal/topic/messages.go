package topic

import (
	"fmt"
	"os"
	"strconv"
	"streaming-queue/internal/infra"
	"streaming-queue/proto"
	"strings"
	"sync"
	"time"
)

var (
	mutex = &sync.Mutex{}
)

type WriteableMessage struct {
	message proto.Message
}

func (w *WriteableMessage) ParseString() string {
	return fmt.Sprintf("%d,%s,%s,%s\n", w.message.Id, w.message.Topic, w.message.Title, w.message.CreatedAt)
}

func ParseMessage(message string) proto.Message {
	fields := strings.Split(message, ",")
	id, topic, title, createdAt := fields[0], fields[1], fields[2], fields[3]

	id32, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return proto.Message{
		Id:        uint32(id32),
		Topic:     topic,
		Title:     title,
		CreatedAt: createdAt,
	}
}

func WriteMessage(message *proto.Message) {
	mutex.Lock()
	defer mutex.Unlock()

	staticFilePath, _ := os.Getwd()
	fileName := fmt.Sprintf("%s/static/%s_messages.txt", staticFilePath, message.Topic)
	infra.WriteOnFile(fileName, &WriteableMessage{message: *message})
}

func ReadMessages(offset, topic string) []proto.Message {
	mutex.Lock()
	defer mutex.Unlock()

	var messages []proto.Message
	staticFilePath, _ := os.Getwd()
	offsetDate, err := time.Parse(time.RFC3339, offset)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return []proto.Message{}
	}

	fileName := fmt.Sprintf("%s/static/%s_messages.txt", staticFilePath, topic)
	infra.ReadFileStream(fileName, func(row string) {
		message := ParseMessage(row)

		messageDate, err := time.Parse(time.RFC3339, message.CreatedAt)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		if messageDate.Before(offsetDate) {
			fmt.Printf("Ignoring message '%s'. Offset date:%s\n", offsetDate.Format(time.DateTime), row)
		} else {
			messages = append(messages, message)
		}
	})

	return messages
}
