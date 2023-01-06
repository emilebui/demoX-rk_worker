package main

import (
	"github.com/emilebui/demoX-rk_worker/internal/handlers"
	"github.com/emilebui/demoX-rk_worker/internal/services"
	"github.com/emilebui/demoX-rk_worker/pkg/config"
	"github.com/emilebui/demoX-rk_worker/pkg/conn"
	"github.com/spf13/viper"
)

func CreateKafkaWorker(addr string, groupid string) handlers.Worker {
	kafkaProcess := services.NewKafkaProcess()
	kafkaConsumer, err := conn.GetKafkaConsumer(addr, groupid, "earliest")
	if err != nil {
		panic(err)
	}
	return handlers.NewKafkaWorker(kafkaConsumer, kafkaProcess)
}

func CreateRedisWorker(conf *viper.Viper) handlers.Worker {
	redisClient := conn.GetRedisConn(conf)
	redisProcess := services.NewRedisProcess(conf, redisClient)
	return handlers.NewRedisWorker(redisProcess)
}

func main() {
	conf := config.Get()

	var worker handlers.Worker

	if conf.GetBool("kafka_mode") {
		worker = CreateKafkaWorker(conf.GetString("kafka_addr"), conf.GetString("kafka_groupid"))
	} else {
		worker = CreateRedisWorker(conf)
	}
	worker.Start()
}
