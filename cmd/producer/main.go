package main

import (
	"context"
	"github.com/mcmacedo/redis-performance-poc/pkg/ports"
	"github.com/mcmacedo/redis-performance-poc/pkg/redis"
	"log"
	"os"
	"strconv"
	"time"
)

func startProducer(ctx context.Context, producer ports.ProducerPort) {
	for {
		_, err := producer.SendMessage(ctx, ports.Message{Data: strconv.FormatInt(time.Now().Unix(), 10)})
		if err != nil {
			log.Fatalf("Producer Finalizado! %v", err)
		}

		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	redisAddress := os.Getenv("REDIS_ADDR") + ":" + os.Getenv("REDIS_PORT")
	redisOptions := redis.Options{Addr: redisAddress}
	redisClient00 := redis.NewClient(redisOptions)

	ctx := context.Background()
	producer00 := redis.NewProducer(os.Getenv("PROD_ID"), redisClient00, os.Getenv("STREAM_NAME"))

	go startProducer(ctx, producer00)

	select {}
}
