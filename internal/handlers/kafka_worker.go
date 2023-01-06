package handlers

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emilebui/demoX-rk_worker/internal/services"
)

type KafkaWorker struct {
	c   *kafka.Consumer
	run bool
	s   *services.KafkaProcess
}

func NewKafkaWorker(c *kafka.Consumer, s *services.KafkaProcess) Worker {
	return &KafkaWorker{
		c:   c,
		run: false,
		s:   s,
	}
}

func (w *KafkaWorker) Start() {

	if w.run {
		fmt.Printf("KafkaWorker already started\n")
		return
	}

	w.run = true
	defer w.c.Close()

	for w.run {

		msg, err := w.c.ReadMessage(-1)
		if err == nil {
			w.s.ProcessTask(msg)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
