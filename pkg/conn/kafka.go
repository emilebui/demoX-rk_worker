package conn

import "github.com/confluentinc/confluent-kafka-go/kafka"

func GetKafkaConsumer(ep string, groupid string, offset string) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}
