package handlers

import (
	"fmt"
	"github.com/emilebui/demoX-rk_worker/internal/services"
	"github.com/go-redis/redis/v9"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RedisWorker struct {
	s       *services.RedisProcess
	run     bool
	sigChan chan os.Signal
}

func NewRedisWorker(s *services.RedisProcess) Worker {
	return &RedisWorker{
		s:       s,
		run:     false,
		sigChan: make(chan os.Signal, 1),
	}
}

func (w *RedisWorker) Start() {

	if w.run {
		fmt.Printf("Redis worker has already started\n")
	}
	w.run = true
	signal.Notify(w.sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	for w.run {

		select {
		case sig := <-w.sigChan:
			fmt.Printf("Process is getting terminated - signal %v\n", sig)
			w.run = false
		default:
			msg, err := w.s.GetMessage()

			if err == redis.Nil {
				continue
			} else if err != nil {
				fmt.Printf("Error getting message: %v\n", err)
				time.Sleep(500 * time.Millisecond)
			} else {
				w.s.ProcessTask(msg)
			}
		}

	}
}
