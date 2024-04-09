package main

import (
	"context"
	"github.com/mcmacedo/redis-performance-poc/pkg/ports"
	"github.com/mcmacedo/redis-performance-poc/pkg/redis"
	"log"
	"os"
)

func startConsumer(ctx context.Context, consumer ports.ConsumerPort) {
	for {
		_, err := consumer.ReceiveMessage(ctx)
		if err != nil {
			log.Fatalf("Consumer Finalizado! %v", err)
		}
	}
}

func main() {
	redisAddress := os.Getenv("REDIS_ADDR") + ":" + os.Getenv("REDIS_PORT")
	redisOptions := redis.Options{Addr: redisAddress}
	redisClient00 := redis.NewClient(redisOptions)

	ctx := context.Background()
	consumer00 := redis.NewConsumer(os.Getenv("CONS_ID"), redisClient00, os.Getenv("STREAM_NAME"))

	go startConsumer(ctx, consumer00)

	select {}
}
