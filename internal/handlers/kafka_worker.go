package handlers

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emilebui/demoX-rk_worker/internal/services"
	"os"
	"os/signal"
	"syscall"
)

type KafkaWorker struct {
	c       *kafka.Consumer
	run     bool
	s       *services.KafkaProcess
	sigChan chan os.Signal
}

func NewKafkaWorker(c *kafka.Consumer, s *services.KafkaProcess) Worker {
	return &KafkaWorker{
		c:       c,
		run:     false,
		s:       s,
		sigChan: make(chan os.Signal, 1),
	}
}

func (w *KafkaWorker) Start() {

	if w.run {
		fmt.Printf("KafkaWorker already started\n")
		return
	}

	w.run = true
	signal.Notify(w.sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	defer w.c.Close()

	for w.run {

		select {

		case sig := <-w.sigChan:
			fmt.Printf("Process is getting terminated - signal %v\n", sig)
			w.run = false
		default:
			ev := w.c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {

			case *kafka.Message:
				w.s.ProcessTask(e)
			case kafka.Error:
				fmt.Printf("%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					w.run = false
				}
			}
		}
	}
}
