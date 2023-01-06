package handlers

import (
	"fmt"
	"github.com/emilebui/demoX-rk_worker/internal/services"
	"time"
)

type RedisWorker struct {
	s   *services.RedisProcess
	run bool
}

func NewRedisWorker(s *services.RedisProcess) Worker {
	return &RedisWorker{
		s:   s,
		run: false,
	}
}

func (w *RedisWorker) Start() {

	if w.run {
		fmt.Printf("Redis worker has already started\n")
	}
	w.run = true

	for w.run {

		msg, err := w.s.GetMessage()
		if err != nil {
			fmt.Printf("Error getting message: %v\n", err)
			time.Sleep(500 * time.Millisecond)
		} else {
			w.s.ProcessTask(msg)
		}
	}
}
