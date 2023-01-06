package services

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type RedisProcess struct {
	rdb  *redis.Client
	conf *viper.Viper
}

func NewRedisProcess(conf *viper.Viper, rdb *redis.Client) *RedisProcess {
	return &RedisProcess{
		conf: conf,
		rdb:  rdb,
	}
}

func (p *RedisProcess) ProcessTask(msg string) {
	fmt.Printf("Message in redis queue: %s\n", msg)
}

func (p *RedisProcess) GetMessage() (string, error) {
	val, err := p.rdb.RPop(context.Background(), p.conf.GetString("redis_queue_key")).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
