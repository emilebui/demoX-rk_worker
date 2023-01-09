package conn

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func GetKafkaConsumer(ep string, groupid string, offset string, topic string) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": ep,
		"group.id":          groupid,
		"auto.offset.reset": offset,
	})
	if err != nil {
		return nil, err
	}
	err = c.Subscribe(topic, nil)
	if err != nil {
		fmt.Printf("Can't subscribe to this topic - %s - The topic might not exist!", topic)
		return nil, err
	}

	return c, nil
}
