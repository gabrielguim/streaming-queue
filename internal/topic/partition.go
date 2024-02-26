package topic

import (
	"streaming-queue/proto"
)

func FindPartition(message *proto.Message, observers *[]Observer) uint32 {
	count := len(*observers)

	partition := message.Id % uint32(count)
	if partition >= uint32(count) {
		partition = uint32(count) - 1
	}

	return partition
}
