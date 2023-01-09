package main

import (
	"bufio"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emilebui/demoX-rk_worker/pkg/config"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func createProducer(conf *viper.Viper) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": conf.GetString("kafka_addr")})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		panic(err)
	}
	return p
}

func pushKafka(p *kafka.Producer, conf *viper.Viper, msg string) {
	topic := conf.GetString("kafka_topic")
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(msg),
	}, nil)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		panic(err)
	}
}

func main() {

	fmt.Println("DemoX Test Producer - Loading Config ...")
	conf := config.Get()

	p := createProducer(conf)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	run := true

	for run {
		select {
		case sig := <-sigChan:
			fmt.Printf("Received signal: %v\n - Exiting the app", sig)
			run = false
		default:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter message to kafka: ")
			text, _ := reader.ReadString('\n')
			pushKafka(p, conf, text)
		}
	}

}
