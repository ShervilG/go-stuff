package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Subscriber struct {
	client *redis.Client
	ctx    context.Context
}

func NewSubscriber(redisAddr string) *Subscriber {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return &Subscriber{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (s *Subscriber) Subscribe(channels ...string) *redis.PubSub {
	return s.client.Subscribe(s.ctx, channels...)
}

func (s *Subscriber) PSubscribe(patterns ...string) *redis.PubSub {
	return s.client.PSubscribe(s.ctx, patterns...)
}
