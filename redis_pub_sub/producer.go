package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Producer struct {
	client *redis.Client
	ctx    context.Context
}

func NewProducer(redisAddr string) *Producer {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		PoolSize: 10,
	})
	return &Producer{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (p *Producer) Publish(channel string, message interface{}) error {
	return p.client.Publish(p.ctx, channel, message).Err()
}
