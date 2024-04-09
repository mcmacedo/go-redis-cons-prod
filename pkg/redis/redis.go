package redis

import "github.com/redis/go-redis/v9"

type Options struct {
	Addr string
}

func NewClient(options Options) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: options.Addr,
	})
}

func NewProducer(id string, client *redis.Client, streamName string) *Producer {
	return &Producer{
		id:         id,
		client:     client,
		streamName: streamName,
	}
}

func NewConsumer(id string, client *redis.Client, streamName string) *Consumer {
	return &Consumer{
		id:         id,
		client:     client,
		streamName: streamName,
	}
}
