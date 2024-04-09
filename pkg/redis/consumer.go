package redis

import (
	"context"
	"errors"
	"github.com/mcmacedo/redis-performance-poc/pkg/ports"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Consumer struct {
	id         string
	client     *redis.Client
	streamName string
}

func (rc *Consumer) ReceiveMessage(ctx context.Context) (ports.Message, error) {
	lastId := "0-0"
	for {
		result, err := rc.client.XRead(ctx, &redis.XReadArgs{
			Streams: []string{rc.streamName, lastId},
			Block:   20 * time.Millisecond,
		}).Result()

		if err != nil {
			if errors.Is(err, redis.Nil) {
				log.Println("Waiting for messages")
				continue
			} else {
				return ports.Message{}, err
			}
		}

		for _, stream := range result {
			for _, message := range stream.Messages {
				lastId = message.ID
				log.Printf("Consumer id: %s Mensagem recebida com ID: %s, Com o info: %s", rc.id, message.ID, message.Values)
			}
		}
	}
}
