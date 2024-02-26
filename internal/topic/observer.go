package topic

import (
	"fmt"
	"slices"
	"streaming-queue/proto"
)

type Observer struct {
	notify func(message *proto.Message)
	id     string
}

type Notifier struct {
	observers *[]Observer
}

func (n *Notifier) Register(observer *Observer) {
	*n.observers = append(*n.observers, *observer)
}

func (n *Notifier) Unregister(observer *Observer) {
	*n.observers = slices.DeleteFunc(*n.observers, func(o Observer) bool {
		return o.id == observer.id
	})
}

func (n *Notifier) Notify(message *proto.Message) {
	count := len(*n.observers)

	if count == 0 {
		fmt.Println("No consumers registered")
		return
	}

	observer := &(*n.observers)[0]
	if count > 1 {
		partition := FindPartition(message, n.observers)

		fmt.Printf("Message going to partition: %d. Active consumers: %d. Target consumer: '%s'\n", partition, count, (*n.observers)[partition].id)
		observer = &(*n.observers)[partition]
	} else {
		fmt.Printf("Message published. Active consumers: %d. Target consumer: '%s'\n", count, observer.id)
	}

	observer.notify(message)
}
