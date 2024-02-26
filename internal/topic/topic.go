package topic

import (
	"fmt"
	"streaming-queue/proto"
)

type Topic struct {
	name     string
	notifier *Notifier
}

func CreateTopic(name string) *Topic {
	return &Topic{
		name,
		&Notifier{
			observers: &[]Observer{},
		},
	}
}

func (t *Topic) RegisterObserver(observer *Observer) {
	t.notifier.Register(observer)
	fmt.Printf("Registering observer '%s' for topic '%s'. Active observers: %d\n", observer.id, t.name, len(*t.notifier.observers))
}

func (t *Topic) UnregisterObserver(observer *Observer) {
	t.notifier.Unregister(observer)
	fmt.Printf("Unregistering observer '%s' for topic '%s'. Active observers: %d\n", observer.id, t.name, len(*t.notifier.observers))
}

func (t *Topic) Publish(message *proto.Message) {
	t.notifier.Notify(message)
}
