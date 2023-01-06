package services

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcess struct {
}

func NewKafkaProcess() *KafkaProcess {
	return &KafkaProcess{}
}

func (p *KafkaProcess) ProcessTask(msg *kafka.Message) {
	// Change logic here
	fmt.Printf("payload: %s from partition %s\n", string(msg.Value), msg.TopicPartition)
}
