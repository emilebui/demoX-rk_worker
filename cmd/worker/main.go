package main

import (
	"fmt"
	"github.com/emilebui/demoX-rk_worker/internal/handlers"
	"github.com/emilebui/demoX-rk_worker/internal/services"
	"github.com/emilebui/demoX-rk_worker/pkg/config"
	"github.com/emilebui/demoX-rk_worker/pkg/conn"
	"github.com/spf13/viper"
)

func CreateKafkaWorker(addr string, groupid string, topic string) handlers.Worker {
	kafkaProcess := services.NewKafkaProcess()
	kafkaConsumer, err := conn.GetKafkaConsumer(addr, groupid, "earliest", topic)
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
	fmt.Println("DemoX RK_WORKER - Loading Config ...")
	conf := config.Get("config.yaml")

	var worker handlers.Worker

	if conf.GetBool("kafka_mode") {

		fmt.Println("Creating kafka worker...")
		worker = CreateKafkaWorker(conf.GetString("kafka_addr"), conf.GetString("kafka_group_id"), conf.GetString("kafka_topic"))
		fmt.Println("Created kafka worker...")
	} else {
		fmt.Println("Creating redis worker...")
		worker = CreateRedisWorker(conf)
		fmt.Println("Created redis worker...")
	}

	fmt.Println("Starting worker...")
	worker.Start()
}
